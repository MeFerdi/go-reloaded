# Project: Text Modification Tool
Welcome to the Text Modification Tool project! This tool allows you to make various modifications to text files based on specific rules. You can perform operations like converting numbers between different bases, changing letter cases, handling punctuation, and more.
### Features

    Convert hexadecimal numbers to decimal numbers
    Convert binary numbers to decimal numbers
    Uppercase, lowercase, and capitalize words
    Handle punctuation placement and formatting
    Manage single quotes around words
    Replace "a" with "an" based on the following word
    Utilize standard Go packages for functionality

### Usage
To use the Text Modification Tool, follow these steps:

    Create a text file with the content that needs modification.
    Run the tool with the input file and the desired output file as arguments.

### Example Usage

bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

Learning Opportunities
By working on this project, you will gain insights into:

    Using the Go file system (fs) API
    Manipulating strings and numbers in Go

### Contribution
To contribute to this project, follow these steps:

    Fork the repository. Fork the repository by clicking the "Fork" button on the top right corner of the project's GitHub page. This will create a copy of the project under your GitHub account.
    Clone the repository. Clone the forked repository to your local machine using git clone command.

    bash
    git clone https://github.com/your-username/text-modification-tool.git

Create a new branch. Create a new branch for your changes using the git checkout command.

bash
git checkout -b your-feature-branch

Replace your-feature-branch with a descriptive name for your changes.
Make your changes. Make your changes in the new branch. Ensure that your changes follow the project's coding standards and conventions.
Test your changes. Test your changes to ensure that they work as expected. You can use the project's existing test suite or create new tests if necessary.
Commit your changes. Commit your changes using the git add and git commit commands.

bash
git add .
git commit -m "Your commit message"

Replace Your commit message with a descriptive message that summarizes your changes.
Push your changes. Push your changes to your forked repository using the git push command.

bash
git push origin your-feature-branch

Replace your-feature-branch with the name of your new branch.
Create a pull request. Create a pull request from your forked repository to the original repository. Provide a detailed description of your changes and why they are necessary.
Review and merge. The project maintainers will review your changes and provide feedback. Once your changes are approved, they will be merged into the main project.
Pull the latest changes. After your changes are merged, pull the latest changes from the main project using the git pull command.

bash
git pull origin main

 Let's dive into the world of text manipulation with the Text Modification Tool! Happy coding! 🚀