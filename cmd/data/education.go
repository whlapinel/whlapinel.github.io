package data

import "gh_static_portfolio/cmd/domain"


type educationRepo struct {
}

func NewEducationRepo() domain.EducationRepository {
	return &educationRepo{}
}

func (r *educationRepo) GetAll() ([]*domain.EducationItem, error) {
	return educationItems, nil
}

var educationItems = []*domain.EducationItem{
	{
		Year:   "2024",
		School: "University of North Carolina at Charlotte",
		Degree: "Master of Science in Information Technology",
		Minor:  "Software Systems Design and Engineering",
		Classes: []*domain.ClassItem{
			{
				Number:      "ITIS 5154",
				Title:       "Applied Machine Learning",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of machine learning, including supervised and unsupervised learning, neural networks, and deep learning.",
			},
			{
				Number:      "ITIS 5153",
				Title:       "Applied Artificial Intelligence",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of artificial intelligence, including expert systems, natural language processing, and computer vision.",
			},
			{
				Number:      "ITIS 6200",
				Title:       "Principles of Information Security and Privacy",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of information security, including cryptography, network security, and risk management.",
			},
			{
				Number:      "ITIS 5180",
				Title:       "Mobile Application Development",
				Subtitle:    "Graduate",
				Description: "This course covers the basics of mobile application development in Android.",
			},
			{
				Number:      "ITIS 6112",
				Title:       "Software Systems Design and Implementation",
				Subtitle:    "Graduate",
				Description: "Introduction to the techniques involved in the planning and implementation of large software systems. Emphasis on human interface aspects of systems. Planning software projects; software design process; top-down design; modular and structured design; management of software projects; testing of software; software documentation; choosing a language for a software system.",
			},
			{
				Number:      "ITIS 6400",
				Title:       "Human-Centered Design",
				Subtitle:    "Graduate",
				Description: "An introduction to Human-Computer Interaction practice and research. Topics include: the perceptual, cognitive, and social characteristics of people, as well as methods for learning more about people and their use of computing systems. The process of interface design, methods of design, and ways to evaluate and improve a design.  Highlights a number of current and cutting-edge research topics in Human-Computer Interaction and is a balance of design, sociological/psychological, and information systems elements.",
			},
			{
				Number:      "ITIS 6177",
				Title:       "Systems Integration",
				Subtitle:    "Graduate",
				Description: "Examines the issues related to system integration. Topics include: data integration, business process integration, integration architecture, middleware, system security, and system management.",
			},
			{
				Number:      "ITIS 6120",
				Title:       "Applied Databases",
				Subtitle:    "Graduate",
				Description: "Identification of business database needs; requirements specification; relational database model; SQL; E-R modeling; database design, implementation, and verification; distributed databases; databases replication; object-oriented databases; data warehouses; OLAP; data mining; security of databases; vendor selection; DBMS product comparison; database project management; tools for database development, integration, and transaction control.",
			},
			{
				Number:      "ITIS 5166",
				Title:       "Network-Based Application Development",
				Subtitle:    "Graduate",
				Description: "Examines the issues related to network based application development. Topics include: introduction to computer networks, web technologies and standards, network based programming methodologies, languages, tools and standards.",
			},
			{
				Number:      "ITIS 5135",
				Title:       "Web-Based Application Development",
				Subtitle:    "Graduate",
				Description: "In this course, students will learn to design and develop interactive webpages with significant focus on programming. Throughout the course, students will develop an interactive and accessible website frontend.",
			},
			{
				Number:      "ITIS 5101",
				Title:       "Foundations of Programming",
				Subtitle:    "Graduate",
				Description: "Foundations of Programming is designed to take the participant from no experience in programming to having an understanding of how to produce software that is relevant to their domain knowledge. Throughout the course, students will learn programming language agnostic, and instruction will focus on making sure that the participant learns to think about problem solutions that are systematic and repeatable (algorithms) and translatable to code.",
			},
		},
	},
	{
		Year:   "2009",
		School: "United States Naval Academy",
		Degree: "Bachelor of Science",
		Minor:  "History",
	},
	{
		Year:   "2001",
		School: "Northwest School of the Arts",
		Degree: "High School Diploma",
		Minor:  "Piano",
	},
}
