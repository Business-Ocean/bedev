# Bedev Terminal App

Bedev is a command-line tool that helps you automate various tasks. It provides a set of commands to simplify common operations in your development workflow.

## Installation

To install Bedev, follow these steps:

1. Clone the repository: `git clone https://github.com/yourusername/bedev`
2. Navigate to the project directory: `cd bedev`
3. Install the dependencies: `npm install`
4. Set up the command globally: `npm link`

## Usage

Bedev provides the following commands:

## Supported Language
 - golang
| - flutter
| - node (todo) 

### 1. `bedev do`
----
This command allows you to perform specific tasks. For example, you can use it to rename files in a folder in a certain format.

Example:
bedev do rename-files --folder path/to/folder --format new-file-{{index}}.txt



### 2. `bedev create`

This command helps you generate various elements, such as avatars or images of a certain size. It can also create gradient or placeholder images.

Example:
bedev create avatar --name John --size 200x200


### 3. `bedev clean`

Use this command to clean your project by removing unnecessary files or folders. For instance, you can clean the Flutter build folder or the Node.js modules folder.

Example:
bedev clean flutter-build



### 4. `bedev download`

This command enables you to download specific files from a given URL. By default, the files will be downloaded to the current path, but you can specify a different destination.

Example:
bedev download https://github.com/yourusername/image.jpg --path path/to/save


### 5. `bedev config`

With this command, you can manage your configuration settings, such as your GitHub user information.

Example:
bedev config set github-username yourusername


### 6. `bedev env`

Use this command to view the environment variables of your operating system in an organized way.

Example:
bedev env


### 7. `bedev convert`

This command allows you to convert files from one format to another. For instance, you can convert PNG images to JPEG format.

Example:
bedev convert image.png --format jpeg



### 8. `bedev generate`

This command allows you to perform generate several thing.
- Generate UUID base of time for user given desire input
- Generate test replication of given directory into specefic test directory (help to cover TDD projects and test cretical software) ex:  route/do_something.go will generate empty test file in test/route/do_something_test.go 


### 8. `bedev run`

This command allows you to perform run script periodically from current directory.


Example:
bedev do rename-files --folder path/to/folder --format new-file-{{index}}.txt


## Contributing

If you want to contribute to Bedev, feel free to fork the repository and submit pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
Feel free to use this consolidated content for your README.md file.



---

# Bedev command structure 

```

1. bedev do  // do some thing for you example: rename file in folder in certain format 

2. bedev create  // create some thing for you example: generate avatar , or image of crertain size , gradiant image placeholder image etc

3. bedev clean // clean project example clean flutter buld folder or node module 

4. bedev download // download only file from certain url example github image or file only will get download to specified path defult will be current path

5. bedev config // setting and getting your configuration like github user info 

6. bedev env // see the environment variable of you os system in organised way

7. bedev convert // convert png into jpeg . etc file


```