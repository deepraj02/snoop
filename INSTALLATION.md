# Installing Snoop :

## MacOS :

Download the Binary from the [Release](https://github.com/deepraj02/snoop/releases)
OR
Use **Homebrew** to install it.

```sh
brew tap deepraj02/taps
```
then 

```sh
brew install snoop
```

**OR**

```sh
brew install deepraj02/taps/snoop
```



## Windows :

> **Note:** Working on to distribute the package with **[Chocolatey](https://chocolatey.org/)** in the coming week.

As for now you can install the zip file for Windows from the [release](https://github.com/deepraj02/snoop/releases) page.

#### Steps : 

- Install the zip for the Release page.
- Unzip and move the snoop.exe file to `C:\Users\<UserName>\AppData\Local\Programs\Snoop\bin`.

    (you may need to create the `snoop/bin` directory)
- Now add the path to your System Variable.

Snoop should work fine now.

## Linux :

#### Steps : 

- Install the `tar.gz` for the [Release](https://github.com/deepraj02/snoop/releases) page.
- Unzip the file and set the path in your `~/.bashrc` file.

> **Note**: If you want to change it permanently add export 
``` bash
PATH=$PATH:</path/to/file>
```
> to your ~/.bashrc file (just at the end is fine).

