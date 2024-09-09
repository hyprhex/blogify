package db

import (
	"context"
	"log"
	"math/rand"

	"github.com/hyprhex/blogify/internal/store"
)

var titles = []string{
	"How to Master Go in 30 Days",
	"The Ultimate Guide to RESTful APIs",
	"10 Tips for Writing Clean Code",
	"Scaling Your Backend with Microservices",
	"Understanding JWT Authentication in Go",
	"Building a Simple Web Server with Go",
	"The Power of Concurrent Programming in Go",
	"Exploring Go's Standard Library",
	"Best Practices for Database Integration in Go",
	"How to Secure Your APIs with OAuth2",
	"Optimizing Go Applications for Performance",
	"Why Go is the Future of Backend Development",
	"Handling Errors Gracefully in Go",
	"Deploying Go Applications with Docker",
	"An Introduction to Go Modules",
	"Creating Reusable Packages in Go",
	"Unit Testing Best Practices in Go",
	"The Benefits of Using Go for Web Development",
	"Understanding Goroutines and Channels in Go",
	"How to Implement Caching in Go Applications",
}

var contents = []string{
	"In this post, we'll explore a structured approach to mastering Go in just 30 days. From syntax basics to advanced concepts like concurrency, you'll gain the skills needed to become proficient.",
	"Learn how to design and implement RESTful APIs using Go. We'll cover everything from setting up your environment to handling requests and responses efficiently.",
	"Clean code is essential for maintainability. Discover 10 practical tips to write cleaner, more readable, and more efficient Go code.",
	"Microservices are a popular architecture for scaling applications. This guide explains how to break down your monolith and scale your backend using Go.",
	"JWT is a powerful tool for securing your APIs. In this post, we'll dive into JWT authentication in Go, with practical examples and security best practices.",
	"Building a web server from scratch in Go is easier than you think. Follow this tutorial to create a simple yet powerful web server that handles HTTP requests.",
	"Concurrency is one of Go's standout features. Learn how to harness the power of Goroutines and Channels to build fast and efficient applications.",
	"Go's standard library is packed with useful tools. This post will introduce you to some hidden gems that can simplify your development process.",
	"Integrating databases into your Go applications can be tricky. We’ll show you best practices for working with popular databases like MySQL and PostgreSQL in Go.",
	"OAuth2 is the industry standard for authorization. Discover how to implement OAuth2 in your Go applications to secure user data and provide seamless login experiences.",
	"Performance matters in production environments. Learn key techniques to optimize your Go applications, from reducing memory usage to speeding up execution.",
	"Go is gaining popularity among backend developers. This post explores why Go is considered the future of backend development and how you can leverage it.",
	"Error handling is crucial in any application. Learn how to handle errors gracefully in Go, ensuring your application remains robust and user-friendly.",
	"Docker makes deploying Go applications simple and consistent. Follow this guide to containerize your Go apps and deploy them anywhere.",
	"Go Modules simplify dependency management. Learn how to create, manage, and publish your own Go Modules with ease.",
	"Reusable packages save time and effort. This post shows you how to design and create reusable Go packages for your projects.",
	"Unit testing ensures code reliability. Discover best practices for writing and organizing unit tests in Go to maintain high code quality.",
	"Go's lightweight nature makes it perfect for web development. Explore the benefits of using Go to build fast, secure, and scalable web applications.",
	"Goroutines and Channels are powerful concurrency tools. In this post, we’ll explain how they work and how to use them effectively in your Go programs.",
	"Caching can significantly improve your application's performance. Learn how to implement caching in Go using popular techniques and libraries.",
}

var categories = []string{
	"Programming",
	"Web Development",
	"Best Practices",
	"Backend",
	"Security",
	"Web Server",
	"Concurrency",
	"Standard Library",
	"Database",
	"Authentication",
	"Performance",
	"Backend Development",
	"Error Handling",
	"Deployment",
	"Dependency Management",
	"Packages",
	"Testing",
	"Web Apps",
	"Concurrency Tools",
	"Caching",
}

var tags = []string{
	"Go, Learning, Programming",
	"Go, RESTful APIs, Web Development",
	"Go, Clean Code, Best Practices",
	"Go, Microservices, Backend",
	"Go, JWT, Authentication",
	"Go, Web Server, HTTP",
	"Go, Concurrency, Goroutines",
	"Go, Standard Library, Tools",
	"Go, Database, MySQL, PostgreSQL",
	"Go, OAuth2, Security",
	"Go, Performance, Optimization",
	"Go, Backend Development, Future Tech",
	"Go, Error Handling, Robustness",
	"Go, Docker, Deployment",
	"Go, Modules, Dependency Management",
	"Go, Reusable Packages, Development",
	"Go, Unit Testing, Code Quality",
	"Go, Web Development, Scalability",
	"Go, Goroutines, Channels, Concurrency",
	"Go, Caching, Performance, Optimization",
}

func Seed(store store.Storage) {
	posts := generatePosts(100)

	ctx := context.Background()

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating posts:", err)
			return
		}
	}

	log.Println("Insert content complete")
}

func generatePosts(num int) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		posts[i] = &store.Post{
			Title:    titles[rand.Intn(len(titles))],
			Content:  contents[rand.Intn(len(contents))],
			Category: categories[rand.Intn(len(categories))],
			Tags:     []string{tags[rand.Intn(len(tags))]},
		}
	}

	return posts
}
