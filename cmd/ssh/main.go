package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/charmbracelet/wish/recover"
	"github.com/muesli/termenv"
	gossh "golang.org/x/crypto/ssh"

	"github.com/daipham3213/mux0.dev/pkg/tui"
	"github.com/google/uuid"
)

const (
	envSSHAddr        = "MUX_SSH_ADDR"
	envHTTPPort       = "MUX_HTTP_PORT"
	envDomain         = "MUX_DOMAIN"
	envSSHHostKeyPath = "MUX_SSH_HOST_KEY_PATH"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	addr := envOrDefault(envSSHAddr, ":22")
	hostKeyPath := envOrDefault(envSSHHostKeyPath, "ssh_host_ed25519_key")
	httpPort := envOrDefault(envHTTPPort, "80")

	srv, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			recover.Middleware(
				bubbletea.Middleware(teaHandler),
				activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
				logging.Middleware(),
			),
		),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			hash := md5.Sum(key.Marshal())
			fingerprint := hex.EncodeToString(hash[:])
			ctx.SetValue("fingerprint", fingerprint)
			ctx.SetValue("anonymous", false)
			return true
		}),
		wish.WithKeyboardInteractiveAuth(
			func(ctx ssh.Context, challenger gossh.KeyboardInteractiveChallenge) bool {
				ctx.SetValue("fingerprint", uuid.NewString())
				ctx.SetValue("anonymous", true)
				return true
			},
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	log.Info("Starting SSH server", "port", addr)
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			cancel()
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	// Listen on port 80
	go func() {
		defer cancel()
		err := http.ListenAndServe(":"+httpPort, nil)
		if err != nil {
			log.Fatal("ListenAndServe error:", err)
		}
	}()

	<-ctx.Done()
	srv.Shutdown(ctx)
	slog.Info("Shutting down server")
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	renderer := lipgloss.NewRenderer(s)
	renderer.SetColorProfile(termenv.EnvColorProfile())
	fingerprint, _ := s.Context().Value("fingerprint").(string)
	anonymous, _ := s.Context().Value("anonymous").(bool)
	command := s.Command()
	slog.Info("got fingerprint", "fingerprint", fingerprint)
	slog.Info("got command", "command", command)

	// Get client IP address from the SSH session
	clientAddr := s.RemoteAddr().String()
	host, _, _ := net.SplitHostPort(clientAddr)
	slog.Info("client connected", "ip", host)

	if pty.Term == "xterm-ghostty" {
		renderer.SetColorProfile(termenv.TrueColor)
	}

	model, err := tui.NewModel(renderer, fingerprint, anonymous, &host, command)
	if err != nil {
		slog.Error("tui init failed", "error", err)
		return nil, []tea.ProgramOption{}
	}
	return model, []tea.ProgramOption{tea.WithAltScreen()}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
