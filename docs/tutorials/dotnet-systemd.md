# Run a .NET Echo Bot as a systemd service on Ubuntu WSL
*Authored by Oliver Smith ([oliver.smith@canonical.com](mailto:oliver.smith@canonical.com))*

In this tutorial we will take advantage of WSL's systemd support to run a chatbot as a systemd service for easier deployment.

We will create the bot using .NET on Ubuntu WSL and it will be accessible from the Windows host.

.NET is an open-source development platform from Microsoft that enables developers to build multi-platform applications from a single codebase.

## Requirements

* A PC running Windows 11
* The latest version of WSL from the Microsoft Store
* Ubuntu, Ubuntu 22.04 LTS or Ubuntu 24.04 LTS
* Visual Studio Code (recommended)

Systemd support is required for this tutorial and is available on WSL version 0.67.6 or higher.

In your PowerShell terminal, you can check your current WSL version by running:

> `wsl --version`

Inside WSL, you can check that systemd is enabled on your Ubuntu distribution with the following command:

> `cat /etc/wsl.conf`

If enabled the output will be:

> `[boot]`  
> `systemd=true`

If systemd is set to `false` then open the file with `sudo nano /etc/wsl.conf`, set it to `true` and save.
Make sure to restart your distribution after you have made this change.

## Install .NET

To install .NET 6 on Ubuntu 24.04 LTS we first need to add the backports archive for .NET.

```{note}
If you are using Ubuntu 22.04 LTS you can skip the command for installing backports and install the .NET 6 bundle directly. 
```

Run this command to install backports, which includes .NET 6:

> `sudo add-apt-repository ppa:dotnet/backports`

To install a bundle with both the SDK and runtime for .NET 6 run:

> `sudo apt install dotnet6`

Run `dotnet --version` to confirm that the package was installed successfully.

## Install and run the Bot Framework EchoBot template

Create a new directory for the project and navigate to it before proceeding:

> `mkdir ~/mybot`  
> `cd mybot`

Once inside we can install the EchoBot C# template by running:

> `dotnet new -i Microsoft.Bot.Framework.CSharp.EchoBot`

We can then verify the template has been installed correctly:

> `dotnet new --list`

You should be able to find the `Bot Framework Echo Bot` template in the list.

![|624x153](assets/dotnet-systemd/templates.png)

Create a new Echo Bot project, with `echoes` as the name for our bot, using the following command:

> `dotnet new echobot -n echoes`

After this has completed we can navigate into the new directory that has been created.

> `cd ~/mybot/echoes`

From inside this directory the project should be ready to run. Test it with:

> `sudo dotnet run`

If everything was set up correctly you should see a similar output to the one below:

![|624x357](assets/dotnet-systemd/welcome-to-dotnet.png)

Leave the EchoBot App running in WSL for now. Open a new browser window on your Windows host and navigate to `localhost:3978` where you should see the following window:

![|624x347](assets/dotnet-systemd/your-bot-is-ready.png)

Leave everything running as we move to the next step.

## Install the Bot Emulator on Windows and connect to your bot

Download the Bot Emulator from the official [Microsoft GitHub](https://github.com/Microsoft/BotFramework-Emulator/releases/tag/v4.14.1), taking care to select [BotFramework-Emulator-4.14.1-windows-setup.exe](https://github.com/microsoft/BotFramework-Emulator/releases/download/v4.14.1/BotFramework-Emulator-4.14.1-windows-setup.exe) and install.

Running it will present you with the following screen, but before you can connect to your bot you need to change a few settings.

![|624x411](assets/dotnet-systemd/bot-framework-emulator.png)

First, get the IP address of your machine by running `ipconfig` in a PowerShell terminal.

![|624x357](assets/dotnet-systemd/ipconfig.png)

Then select the setting icon in the bottom-left corner of the Bot Framework Emulator and enter your IP under ‘localhost override'.

![|624x411](assets/dotnet-systemd/emulator-settings.png)

Click 'Save' and navigate back to the Welcome tab.

Click ‘Open Bot’ and under ‘Bot URL’ input:

```text
http://localhost:3978/api/messages
```

![|624x411](assets/dotnet-systemd/open-a-bot.png)

And click ‘Connect’ to connect to your Echo Bot running in WSL and start chatting!

![|624x411](assets/dotnet-systemd/start-chatting.png)

Congratulations, your Echo Chat Bot App is running on Ubuntu WSL as an App. Now it is time to make it run as a service.

## Running your Echo Bot as a systemd service

Return to your running WSL distro and end the app with `Ctrl+C`.

Then install the .NET systemd extension with:

> `sudo dotnet add package Microsoft.Extensions.Hosting.Systemd`

We can open our project with VS Code by running this command in the 'echoes' directory:

> `code .`

Navigate to ‘Program.cs’ and insert `.UseSystemd()` as a new line in the location shown in the screenshot.

![|624x381](assets/dotnet-systemd/program-cs.png)

Save and close the project in VS Code and return to your WSL terminal.

Next we need to create a service file for your bot using your favourite editor, for example.

> `sudo nano /etc/systemd/system/echoes.service`

Then paste the snippet below taking care to replace `<your-username>` with your username.

```text
[Unit]
Description=The first ever WSL Ubuntu systemd .NET ChatBot Service

[Service]
WorkingDirectory=/home/<your-username>/mybot/echoes
Environment=DOTNET_CLI_HOME=/temp
ExecStart=dotnet run
SyslogIdentifier=echoes

[Install]
WantedBy=multi-user.target
```

![|624x401](assets/dotnet-systemd/nano-service-file.png)

Save your file and reload your services with:

> `sudo systemctl daemon-reload`

To reload the services. You can check if your service is ready by running:

> `systemctl status echoes.service`

You should get the following output:

![|624x77](assets/dotnet-systemd/systemctl-status-inactive.png)

Now start your service and then check its status again.

> `sudo systemctl start echoes.service`

> `sudo systemctl status echoes.service`

If everything has been configured correctly you should get an output similar to the below.

![|624x401](assets/dotnet-systemd/systemctl-status-running.png)

Return to your Windows host and reconnect to your Bot Emulator using the same information as before and confirm that your bot is running, but this time as a systemd service!

![|624x411](assets/dotnet-systemd/start-chatting-service.png)

You can stop your bot from running at any time with the command:

> `sudo systemctl stop echoes.service`

## Tutorial complete!

You now have a simple Echo Bot running as a systemd service on WSL that you can access from your host Windows machine.

If you would like to expand on this example try reviewing some of the more advanced [Bot Framework samples](https://github.com/Microsoft/BotBuilder-Samples/blob/main/README.md) on the Microsoft GitHub.

To read more about how Ubuntu supports .NET developers, making it easier than ever to build multi-platform services and applications, read our [previous announcement](https://ubuntu.com/blog/install-dotnet-on-ubuntu).

### Further Reading

* [.NET on Ubuntu](https://ubuntu.com/blog/install-dotnet-on-ubuntu)
* [Bot Framework samples](https://github.com/Microsoft/BotBuilder-Samples/blob/main/README.md)
* [Working with Visual Studio Code on Ubuntu WSL](vscode.md)
* [Enabling GPU acceleration on Ubuntu on WSL2 with the NVIDIA CUDA Platform](gpu-cuda.md)
* [Windows and Ubuntu interoperability on WSL2](interop.md)
* [Microsoft WSL Documentation](https://learn.microsoft.com/en-us/windows/wsl/)
* [Ask Ubuntu](https://askubuntu.com/)
