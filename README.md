LinkedIn Automation – Technical Proof of Concept (Go + Rod)
Disclaimer

This project is created strictly for educational and technical evaluation purposes.
Automating LinkedIn violates LinkedIn’s Terms of Service.
This tool must never be used on real LinkedIn accounts or deployed in production.

The repository exists only to demonstrate browser automation concepts, anti-detection techniques, and clean Go architecture in a controlled proof-of-concept environment.

Project Overview

This repository contains a Go-based browser automation proof of concept built using the Rod library.

I built this project to demonstrate my ability to quickly learn a new language (Golang), understand browser automation fundamentals, implement human-like behavior patterns, and work within ethical, legal, and security constraints.

The primary focus of this assignment is automation architecture and stealth engineering, not building a real-world LinkedIn automation tool.

Objectives

The objectives of this project are:

Demonstrate advanced browser automation using Go

Implement realistic human-like interaction patterns

Apply multiple anti-detection (stealth) techniques

Maintain clean, modular, and maintainable code

Respect ethical boundaries and platform policies

Implemented Stealth Techniques

The following human-based and anti-detection techniques are implemented as part of this proof of concept:

Randomized timing patterns between actions

Human-like scrolling behavior with pauses and direction changes

Realistic typing simulation with character-by-character input

Mouse hover behavior before clicking elements

Business-hour execution scheduling

Rate limiting and throttling to cap daily actions

State persistence using local JSON storage

Browser fingerprint masking using a realistic User-Agent

These techniques demonstrate an understanding of how automation can be made to behave more like real human interaction.

Core Functional Modules (Proof of Concept)

This project includes structured proof-of-concept modules representing the core functional requirements described in the assignment.

Authentication System (Proof of Concept)

Login using credentials from environment variables (documented)

Human-like login behavior

Login failure handling (conceptual)

Security checkpoint detection such as 2FA and captcha (conceptual)

Session cookie persistence (conceptual)

Search and Targeting (Proof of Concept)

Search users by job title, company, location, and keywords

Parse and collect profile URLs

Handle pagination across search results

Detect and avoid duplicate profiles

Connection Requests (Proof of Concept)

Navigate to user profiles programmatically

Click the Connect button with precise targeting

Send personalized connection notes within limits

Track sent requests and enforce daily limits

Messaging System (Proof of Concept)

Detect newly accepted connections

Send follow-up messages automatically

Support templates with dynamic variables

Maintain message tracking

Important Note:
High-risk actions such as real authentication, connection requests, and messaging are intentionally represented as proof-of-concept modules rather than executed on live LinkedIn accounts, in accordance with the assignment’s educational disclaimer.

Project Structure

The project follows a clean and modular structure with separate packages for the main application, stealth logic, state persistence, authentication, search, connections, messaging, and configuration management.

Configuration

This project uses environment-based configuration.
A sample environment file named .env.example is included to demonstrate required variables such as email, password, daily limits, and user agent values.
No real credentials are required or used.

How to Run

Prerequisites:

Go version 1.20 or higher

Google Chrome installed

Windows operating system

Steps:

Clone the repository

Navigate to the project directory

Run the application using the Go command

When executed, the application opens a visible Chrome browser and demonstrates controlled, human-like automation behavior.

Challenges and Learnings

During this assignment, I learned how to:

Build browser automation systems in Go

Handle real-world OS constraints such as Windows Defender

Synchronize browser actions with realistic timing

Design modular automation architectures

Implement ethical and safe automation proofs of concept

Technologies Used

Go programming language

Rod browser automation library

Google Chrome

JSON for state persistence

AI-assisted tools for learning and debugging

Proof of Concept Scope

This repository represents a technical proof of concept, not a production-ready automation tool.

The focus is on stealth mechanisms, automation architecture, and learning speed rather than real-world LinkedIn automation.

Author

Achitha Vara

Final Note

This project demonstrates my understanding of browser automation, stealth techniques, and system design while respecting platform policies and legal boundaries.