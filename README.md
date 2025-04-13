# Learnado: Igniting Minds, Inspiring Learning (Student Edition)

## Welcome to Learnado!

Learnado is an offline-first learning platform designed to provide you with access to quality educational content without the need for a constant internet connection. Our platform empowers you to learn at your own pace, anytime, anywhere, making education accessible and convenient.

## Getting Started

To start your learning journey with Learnado, follow these simple steps:

### 1. Downloading Learnado

Before you can start using Learnado, you need to download the Learnado software. Clone the repository to get started.

Once you have downloaded and extracted the Learnado package, you can proceed with the remaining steps.

### 2. Configure Environment Variables

Using a text editor, open the `variables.env` file for editing.

Locate the line in the `variables.env` file that begins with `URL=`. This line specifies the server URL for the content manager edition of Learnado.

Replace the existing value after the "=" sign with the server URL provided by your educational institution or content provider. Make sure to enter the complete URL, including the protocol (http://).

Save the changes to the variables.env file.

### 3. Executing the Learnado Binary

To run Learnado on your system, you need to execute the Learnado binary. Here are the steps to execute the binary on different operating systems:

#### For MacOS and Linux:
1. Open a terminal.
2. Navigate to the directory where the Learnado binary is located.
3. Use the `cd` command followed by the directory path to navigate to the Learnado folder. For example:
   ```
   cd /path/to/learnado
   ```
4. Once inside the Learnado folder, run the Learnado binary by executing the following command:
   ```
   go run main.go
   ```
   Note: If you encounter permission issues, you may need to make the Learnado binary executable by running the command `chmod +x Learnado-Student` before executing it.

#### For Windows:
1. Open the command prompt or PowerShell.
2. Navigate to the directory where the Learnado binary is located. This may be the folder where you downloaded or cloned the Learnado repository.
3. Use the `cd` command followed by the directory path to navigate to the Learnado folder. For example:
   ```
   cd C:\path\to\learnado
   ```
4. Once inside the Learnado folder, run the Learnado binary by executing the following command:
   ```
   go run main.go
   ```

After executing the Learnado binary, the Learnado platform will start running on your local machine. If it successfully runs, you will see this message: `⇨ http server started on [::]:3000`.

### 4. Activation

You will receive a license key from your educational institution or content provider. Activate your Learnado software using this license key to gain access to the courses you are authorized to download.

To activate your software, go to [http://localhost:3000/register/xxxx-xxxx-xxxx](http://localhost:3000/register/xxxx-xxxx-xxxx) (replace "xxxx-xxxx-xxxx" with your actual license key) in your web browser.

If your license activation is successful, you will see this message: `License registered`.

### 5. Learning Offline

Once the courses are downloaded, you can dive into your learning journey without worrying about internet access. 

Go to [http://localhost:3000](http://localhost:3000) to access your course content.

Learnado provides a user-friendly interface for seamless navigation.

### 6. Periodic Updates

Learnado periodically checks for new course content and updates. When connected to the internet, it automatically downloads any available updates, ensuring that you always have access to the latest educational materials.

## Key Features

### Offline Learning

Learnado is specifically designed for offline use, making it ideal for areas with limited or unreliable internet connectivity. Once the courses are downloaded, you can access them anytime, even without an internet connection.

### Self-Paced Learning

Learnado supports self-paced learning, allowing you to progress through the course materials at your own speed. You have the flexibility to revisit topics, pause and resume your learning whenever you want, and track your progress along the way.

### Secure Content Delivery

Your course content is compressed and encrypted to ensure its protection against piracy. Learnado's robust security measures safeguard your access to high-quality educational materials.

### Easy Content Updates

Learnado periodically checks for new course content and updates, allowing you to stay up-to-date with the latest educational resources. Simply connect to the internet to download the updates and expand your learning opportunities.

## System Requirements

Learnado is designed to run on various devices and operating systems, including Windows, Linux, and MacOS.

Thank you for choosing Learnado as your learning companion. Start exploring the world of knowledge and unlock your learning potential with Learnado today!
