# Go AMI Event Listener

A Go-based client for the **Asterisk Manager Interface (AMI)**. This tool connects to your Asterisk PBX, authenticates, and streams real-time events to both your console and a persistent log file.


## Prerequisites

* **Go**: Version 1.18 or higher.
* **Asterisk**: An active server with AMI enabled in `manager.conf`.
* **Permissions**: Ensure your AMI user has `read` privileges for the events you want to capture.

## Installation

1. **Clone the repository**:
   ```bash
   git clone [https://github.com/yourusername/your-repo-name.git](https://github.com/yourusername/your-repo-name.git)
   cd your-repo-name
