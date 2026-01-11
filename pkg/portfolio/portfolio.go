package portfolio

type Section struct {
	ID    string
	Title string
	Body  string
}

type Portfolio struct {
	Name     string
	Sections []Section
}

func Default() Portfolio {
	return Portfolio{
		Name: "Pham Le Gia Dai",
		Sections: []Section{
			{
				ID:    "about",
				Title: "About",
				Body: `Hi there! I am Pham Le Gia Dai and welcome to my portfolio.

I am a software engineer with a passion for building things. I have experience
in web development, cloud computing, and DevOps. I am always looking for new
challenges and opportunities to learn and grow.

This project is a TUI app that showcases my portfolio. You can use it to learn
more about me, my work experience, education, certificates, and social media
links.

You can also check out the source code at https://github.com/daipham3213/mux0.dev
or reach out to me at daipham.3213@gmail.com if you want to say hi.`,
			},
			{
				ID:    "experience",
				Title: "Experience",
				Body: `Fullstack Developer | FPT Smart Cloud | 2022 - present

FPT File Storage
- Designed and developed a high-availability File Storage service with a multi-tier storage architecture,
  optimizing performance and cost efficiency for diverse workloads.
- Implemented a high-performance storage tier tailored for AI and machine learning workloads, reducing
  data retrieval times and accelerating training processes.
- Enabled seamless scalability and reliability, ensuring data availability and integrity under high demand.
- Integrated with OpenStack and other open-source technologies to provide a flexible and vendor-agnostic
  storage solution.
- Tech stack: Python, ReactJS, VAST, DDN, Prometheus, Grafana

FPT VPNaaS
- Designed and developed a new VPN as a Service (VPNaaS) solution compatible with OpenStack, enabling
  seamless, scalable, and highly available networking for users.
- Integrated an exporter for the monitoring service, improving system observability and reducing incident
  response time by 40%.
- Developed management tools for VPNaaS, integrating with the OpenStack CLI to streamline operations and
  reduce manual configuration time by 50%.
- Supported the migration and configuration of VPNaaS, ensuring a smooth transition with zero downtime
  and improved network reliability.
- Tech stack: Python, Prometheus, Grafana, RabbitMQ, OpenStack (Neutron, Cinder, Keystone, Glance, Nova),
  OpenStack CLI, K8s, ArgoCD.

FPT Object Storage
- Refactored and enhanced the existing codebase, improving maintainability and reducing technical debt,
  leading to a 25% increase in development efficiency.
- Added support for multi-backend Software-Defined Storage (SDS), enhancing flexibility and optimizing
  storage performance across diverse workloads.
- Improved the UX/UI of the Object Storage service on the FPT Cloud portal, streamlining user
  interactions and boosting user satisfaction by 30%.
- Tech stack: Python, ReactJS, CEPH.

FPT Backup Native
- Developed an agent-less backup solution for FPT Cloud Platform.
- Built an exporter for monitoring service.
- Integrated with the billing service for FPT Cloud Platform.
- Tech stack: Python, Prometheus, Grafana, Kafka, RabbitMQ, OpenStack Cinder, AWX, Ansible, K8s.

FPT AutoScaling
- Integrated OpenStack Senlin with FPT Cloud Platform.
- Developed a new driver for integrating Senlin with VMWare vCenter.
- Developed a new auto-scaling policy supporting AVI load balancer.
- Tech stack: Python, OpenStack Senlin, VMWare, AVI load balancer.

Intern Backend Developer | CMC TSSG | 2021 - 2022
- Learned about the software development process.
- Tech stack: Django, DRF, MySQL, Docker, Git.`,
			},
			{
				ID:    "education",
				Title: "Education",
				Body: `Ho Chi Minh City University of Technology
Software Engineering | 2018 - 2022
GPA: 3.32/4.0`,
			},
			{
				ID:    "awards",
				Title: "Certificates & Awards",
				Body: `Certificates
- Certified Kubernetes Administrator (CKA) | 2025
- Certified Kubernetes Application Developer (CKAD) | 2024

Nominations
- OpenStack Caracal Contributor | https://www.openstack.org/software/openstack-caracal/
- OpenStack Dalmatian Contributor | https://www.openstack.org/software/openstack-dalmatian-2/

Meetups
- Distributed tracing for microservices in OpenStack | https://superuser.openinfra.org/articles/viet-openinfra-user-group-holds-its-30th-meetup/`,
			},
			{
				ID:    "social",
				Title: "Socials",
				Body: `GitHub   | https://github.com/daipham3213
Launchpad | https://launchpad.net/~daiplg
Gerrit   | https://review.opendev.org/q/owner:daiplg
LinkedIn | https://www.linkedin.com/in/daipham-3213
Email    | daipham.3213@gmail.com

Timezone: UTC+07 (Asia/Ho_Chi_Minh)
Please feel free to reach out to me!`,
			},
		},
	}
}
