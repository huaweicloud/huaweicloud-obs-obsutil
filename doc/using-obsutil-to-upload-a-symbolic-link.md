# Using obsutil to Upload a Symbolic Link<a name="EN-US_TOPIC_0177232037"></a>

obsutil supports the upload of the real path to which the symbolic link points when a file or folder is uploaded. You can specify the command-level parameter  **link**  to implement this function when running commands for upload or incremental synchronization upload.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   obsutil can identify symbolic links pointing to folders. If a symbolic link points to a folder, obsutil recursively scans the contents in the folder.  
>-   Avoid the symbolic link loop of a folder, otherwise, the upload will exit due to panic. If you do not want the system to panic, set  **panicForSymbolicLinkCircle**  to  **false**  in the configuration file.  
>-   The symbolic link and the shortcut on the Windows OS are two different types. obsutil cannot identify the shortcut on the Windows OS.  

