# Jobhunt Automation Message Project


### Overview

An automation tool for creating messages during a job hunt. Just setup your configuration in a JSON file, write a message using Golang's `template` syntax, and the tool will interpolate values from your configuration file into `result.txt`.

### Requirements

* Golang `1.9.x`
* a text editor

### Usage

Write your configuration file in `./config.json` with the following shape:

```
{
  "User": {
    "Role": {
      "Name": "Content Specialist",
      "Duties": "writing SQL queries for data engineers"
    },
    "Name": "Dev Developerson"
  },
  "Open": {
    "Role": {
      "Name": "SQL Developer",
      "Duties": "write SQL queries and administer a database"
    },
    "Company": {
      "Name": "Aces, Inc.",
      "Goals": "developing robust solutions for non-trivial problems in the security space"
    }
  }
}
```

Write your message in a `./message.txt` file using Golang's [template|(https://golang.org/pkg/text/template/)] syntax:

```
Greetings! My name is {{.User.Name}}, and I'm an experienced {{.User.Role.Name}}. I noticed that you posted an opening for the position of {{.Open.Role.Name}}. Your company would benefit from leveraging my experience in {{.User.Role.Duties}}. Specifically, the company would be able to {{.Open.Company.Goals}} through the output that a {{.Open.Role.Name}} with the knowledge I've gained over time as a {{.User.Role.Name}} would produce. You'll find that, more than the existence of a good fit, that there is much to be gained from having my knowledge in {{.User.Role.Duties}}.

We can eliminate any doubts about what the company stands to gain by having a quick chat about my qualifications in slightly greater detail.

Best,
```

Run `go run main.go` to run the tool and generate a `result.txt` file in the `build` folder. This file will contain your final message.
