# Setting obsutil Commands to Built-in Commands<a name="EN-US_TOPIC_0150341074"></a>

## Scenario<a name="section0564832202219"></a>

Because obsutil is external software, you need to access the directory where obsutil resides before running obsutil commands. In this way, the usability of the tool is poor.

An OS provides built-in commands so that directories which support running of the commands are loaded to the memory when the system is started. In this way, you can run commands in any directory, which improves the tool's usability.

This section introduces how to set obsutil commands to built-in commands in different OSs.

## Setting obsutil Commands to Built-in Commands in Windows<a name="section168751940102118"></a>

1.  In the CLI, run the  **echo %PATH%**  command to query all the paths configured in the current system. Then select one as the operation path.
2.  Run the  **mklink **_PATH_**/obsutil.exe **_OBSUTIL\_PATH_  command to set obsutil commands to built-in commands of the system.

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >_PATH_  indicates the operation path selected in step 1.  _OBSUTIL\_PATH_  indicates the absolute path of  **obsutil.exe**.  

3.  Check whether the configuration is successful: Run the  **obsutil help**  command in the CLI. If the help information is displayed, the configuration is successful.

## Setting obsutil Commands to Built-in Commands in Linux or macOS<a name="section524910267228"></a>

1.  Run the following command to create a directory for the obsutil tool:

    ```
    mkdir /obsutil
    ```

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >-   Skip this step if the directory already exists.  
    >-   You must run the command as user  **root**.  

2.  Run the following command to grant the  **755**  permission for the tool's directory:

    ```
    chmod 755 /obsutil
    ```

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >-   Skip this step if the permission for the directory is  **drwxr-xr-x**.  
    >-   You must run the command as user  **root**.  

3.  Copy the obsutil tool to the directory created in step 1 and change its permission to  **711**. Assume that the original path of the tool is  **/home/test/obsutil**. Run the following command:

    ```
    cp /home/test/obsutil /obsutil
    chmod 711 /obsutil/obsutil
    ```

4.  Run the  **vi /etc/profile**  command, type  **I**  to enter the Insert mode to edit the file. Add  **export PATH=$PATH:/obsutil**  at the end of the file. Then press  **ESC**  to exit the editing mode, and then type  **:wq!**  and press  **Enter**  to save the file and exit.

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >Skip this step if the new line already exists in the  **/etc/profile**  file.  

5.  Run the  **echo $PATH**  command to query the current environment variables. If  **:/obsutil**  in included in the query result, indicating that the  **/obsutil**  environment variable already exists, go to the next step. Otherwise, run the  **source /etc/profile**  command.
6.  Check whether the configuration is successful: Run the  **obsutil help**  command in any directory. If the help information is displayed, the configuration is successful.

## FAQs<a name="section14418131348"></a>

1.  How do I locate the obsutil configuration file after setting obsutil commands to built-in commands?

    The  **.obsutilconfig**  file in the same directory where obsutil commands reside is the configuration file of the obsutil tool. You can also run the  **obsutil config**  command to obtain the configuration file path. An example is provided as follows:

    ```
    obsutil config
    Config file url:
      D:\tools\.obsutilconfig
    ```

2.  How do I delete obsutil commands after setting them as built-in commands?
    -   In Windows:
        1.  Run the  **where obsutil**  command to locate the path of obsutil commands.

            ```
            where obsutil
            E:\tools\bin\obsutil.exe
            ```

        2.  Run the  **del **_PATH_  command to delete obsutil commands.

            ```
            del E:\tools\bin\obsutil.exe
            ```

            >![](public_sys-resources/icon-note.gif) **NOTE:**   
            >Replace  _PATH_  with the path of obsutil commands.  **E:\\tools\\bin\\obsutil.exe**  is used in the preceding example.  


    -   In Linux or macOS:
        1.  Run the  **which obsutil**  command to locate the path of obsutil commands.

            ```
            which obsutil
            /obsutil/obsutil
            ```

        2.  Run the  **rm -rf **_PATH_  command to delete obsutil commands.

            ```
            rm -rf /obsutil/obsutil
            ```

            >![](public_sys-resources/icon-note.gif) **NOTE:**   
            >Replace  _PATH_  with the path of obsutil commands.  **/obsutil/obsutil**  is used in the preceding example.  

        3.  Restore the system environment variable: Delete the path of obsutil that is set in the  **/etc/profile**  file.

            >![](public_sys-resources/icon-note.gif) **NOTE:**   
            >If the  **/etc/profile**  file contains line  **export PATH=$PATH:/obsutil**, delete the line. Or if the file contains line  **export PATH=$PATH:/test/bin:/obsutil:/test1**, delete  **:/obsutil**  from the line.  


3.  What should I do if the execution of built-in obsutil commands fails in Linux or macOS?
    -   If the message  **Permission denied**  is displayed after executing  **obsutil help**, run the  **chmod 755 **_OBSUTIL\_PATH_  command \(replace  _OBSUTIL\_PATH_  with the path of obsutil\) to add an execute permission for the obsutil tool.
    -   If the message  **command not found**  is displayed, log in again.
    -   If the message  **Cannot create parent folder for xx/.obsutilconfig, xx Permission denied**  is displayed, check whether the home directory of the user exists.

        >![](public_sys-resources/icon-notice.gif) **NOTICE:**   
        >In the Ubuntu OS, if you run the  **useradd**  command to add a user, the home directory of the user is not created by default. You need to create it manually. Therefore, you are advised to run the  **adduser**  command to add a user.  


4.  What can I do if no log file is generated after running built-in obsutil commands in Linux or macOS?

    If you have properly configured  **sdkLogPath**  and  **utilLogPath**  in the configuration file, but still no log file is generated after command execution, then check whether the user who runs the command has the read and write permissions on  **sdkLogPath**  and  **utilLogPath**.


