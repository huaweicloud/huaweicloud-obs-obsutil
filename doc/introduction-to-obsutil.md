# Introduction to obsutil<a name="EN-US_TOPIC_0142009727"></a>

obsutil is a command line tool for accessing and managing OBS. You can use this tool to perform common configurations in OBS, such as creating buckets, uploading and downloading files/folders, and deleting files/folders. If you are familiar with command line interface \(CLI\), obsutil is recommended as an optimal tool for batch processing and automated tasks.

obsutil is compatible with the Windows, Linux, and macOS operating systems \(OSs\).  [Table 1](#table1282175018104)  lists the recommended OS versions. 

**Table  1**  Recommended OS versions for using obsutil

<a name="table1282175018104"></a>
<table><thead align="left"><tr id="row1083550111016"><th class="cellrowborder" valign="top" width="22.97%" id="mcps1.2.3.1.1"><p id="p78385041015"><a name="p78385041015"></a><a name="p78385041015"></a>OS</p>
</th>
<th class="cellrowborder" valign="top" width="77.03%" id="mcps1.2.3.1.2"><p id="p14841250121019"><a name="p14841250121019"></a><a name="p14841250121019"></a>Recommended Version</p>
</th>
</tr>
</thead>
<tbody><tr id="row17841250151013"><td class="cellrowborder" valign="top" width="22.97%" headers="mcps1.2.3.1.1 "><p id="p984850121013"><a name="p984850121013"></a><a name="p984850121013"></a>Windows</p>
</td>
<td class="cellrowborder" valign="top" width="77.03%" headers="mcps1.2.3.1.2 "><a name="ul1569975214153"></a><a name="ul1569975214153"></a><ul id="ul1569975214153"><li>Windows 7</li><li>Windows 8</li><li>Windows 10</li><li>Windows Server 2016</li></ul>
</td>
</tr>
<tr id="row13841150191018"><td class="cellrowborder" valign="top" width="22.97%" headers="mcps1.2.3.1.1 "><p id="p88475041016"><a name="p88475041016"></a><a name="p88475041016"></a>Linux</p>
</td>
<td class="cellrowborder" valign="top" width="77.03%" headers="mcps1.2.3.1.2 "><a name="ul9872184517166"></a><a name="ul9872184517166"></a><ul id="ul9872184517166"><li>SUSE 11</li><li>EulerOS 2</li></ul>
</td>
</tr>
<tr id="row1384185091013"><td class="cellrowborder" valign="top" width="22.97%" headers="mcps1.2.3.1.1 "><p id="p158410500108"><a name="p158410500108"></a><a name="p158410500108"></a>macOS</p>
</td>
<td class="cellrowborder" valign="top" width="77.03%" headers="mcps1.2.3.1.2 "><p id="p17842505108"><a name="p17842505108"></a><a name="p17842505108"></a>macOS 10.13.4</p>
</td>
</tr>
</tbody>
</table>

## Open-Source Addresses<a name="section9338140135813"></a>

-   Click  [here](https://github.com/huaweicloud/huaweicloud-obs-obsutil)  to download the source code of the latest version.
-   Click  [here](https://github.com/huaweicloud/huaweicloud-obs-obsutil/tree/master/release)  to download the compiled software packages of historical versions.

## Tool Advantages<a name="section1963918332594"></a>

obsutil features the following aspects:

1.  Simple and easy to use
2.  Lightweight and installation-free
3.  Compatible with Windows, Linux, and macOS operating systems
4.  Diversified configurations and excellent performance

## Application Scenarios<a name="section38975716584"></a>

-   Automated backup and archiving, for example, periodically uploading local data to OBS.
-   Scenarios that cannot be implemented using other tools such as OBS Browser, for example, synchronously uploading, downloading, and copying objects.

## Functions<a name="section3951253155910"></a>

[Table 2](#table1233162315227)  lists obsutil functions.

**Table  2**  obsutil functions

<a name="table1233162315227"></a>
<table><thead align="left"><tr id="row034112362217"><th class="cellrowborder" valign="top" width="27.51%" id="mcps1.2.3.1.1"><p id="p1734192382220"><a name="p1734192382220"></a><a name="p1734192382220"></a><strong id="b1711191919312"><a name="b1711191919312"></a><a name="b1711191919312"></a>Function</strong></p>
</th>
<th class="cellrowborder" valign="top" width="72.49%" id="mcps1.2.3.1.2"><p id="p134162392219"><a name="p134162392219"></a><a name="p134162392219"></a><strong id="b95394221138"><a name="b95394221138"></a><a name="b95394221138"></a>Description</strong></p>
</th>
</tr>
</thead>
<tbody><tr id="row1934162372212"><td class="cellrowborder" valign="top" width="27.51%" headers="mcps1.2.3.1.1 "><p id="p53462315226"><a name="p53462315226"></a><a name="p53462315226"></a><a href="bucket-commands.md">Basic operations on buckets</a></p>
</td>
<td class="cellrowborder" valign="top" width="72.49%" headers="mcps1.2.3.1.2 "><p id="p43462318225"><a name="p43462318225"></a><a name="p43462318225"></a>Create buckets of different storage classes in specific regions, delete buckets, and obtain the bucket list and configuration information.</p>
</td>
</tr>
<tr id="row33415236221"><td class="cellrowborder" valign="top" width="27.51%" headers="mcps1.2.3.1.1 "><p id="p4341423112213"><a name="p4341423112213"></a><a name="p4341423112213"></a><a href="object-commands.md">Basic operations on objects</a></p>
</td>
<td class="cellrowborder" valign="top" width="72.49%" headers="mcps1.2.3.1.2 "><p id="p15341923152219"><a name="p15341923152219"></a><a name="p15341923152219"></a>Manage objects, including uploading, downloading, deleting, and listing objects. Supported operations are detailed as follows:</p>
<a name="ul12642182419529"></a><a name="ul12642182419529"></a><ul id="ul12642182419529"><li>Upload one or more files or folders.</li><li>Upload large files in multiple parts.</li><li>Synchronously upload, download, and copy incremental objects.</li><li>Copy a single object or copy multiple objects in batches by object name prefix.</li><li>Move a single object or move objects in batches by object name prefix.</li><li>Resume failed upload, download, or copy tasks.</li></ul>
</td>
</tr>
<tr id="row181665745717"><td class="cellrowborder" valign="top" width="27.51%" headers="mcps1.2.3.1.1 "><p id="p13166074572"><a name="p13166074572"></a><a name="p13166074572"></a><a href="log-files.md">Logging</a></p>
</td>
<td class="cellrowborder" valign="top" width="72.49%" headers="mcps1.2.3.1.2 "><p id="p616618717570"><a name="p616618717570"></a><a name="p616618717570"></a>Allows you to configure logging on the client side to record operations on buckets and objects for statistics analysis later.</p>
</td>
</tr>
</tbody>
</table>

## Command Line Structure<a name="section547113329618"></a>

The obsutil command line structures are as follows:

-   In Windows

    **obsutil **_command_ _\[parameters...\]_ _\[options...\]_

-   In Linux or macOS

    ****./**obsutil **_command \[parameters...\] \[options...\]_


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   **command**  indicates the command to be executed, for example,  **ls**  or  **cp**.  
>-   **parameters**  indicates the basic parameters \(mandatory\) of the command, for example, bucket name when creating a bucket.  
>-   **options**  indicates the additional parameters \(optional\) of the command. Additional parameters must be preceded with a hyphen \(-\) when you run the command.  
>-   The square brackets \(\[\]\) are not part of the command. Do not enclose parameter values with them when entering a command.  
>-   If the command contains special characters including ampersands \(&\), angle brackets \(<\) and \(\>\), and spaces, they need to be escaped using quotation marks. Use single quotation marks for Linux or macOS and quotation marks for Windows.  
>-   Additional parameters can be input in the  **-**_key_**=**_value_  or  **-**_key_** **_value_  format, for example,  **-acl=private**, or  **-acl private**. There is no difference between the two formats. Select either one as you like.  
>-   In Windows, you can directly execute  **obsutil.exe**  to enter an interactive command mode. In this mode, you can input  _command \[parameters...\] \[options...\]_  without  **obsutil**  to run a command. An example is provided as follows:  
>    ```  
>    Enter "exit" or "quit" to logout  
>    Enter "help" or "help command" to show help docs  
>    Input your command:  
>    -->ls -limit=3 -s  
>    obs://bucket-001  
>    obs://bucket-002  
>    obs://bucket-003  
>    Bucket number is: 3  
>    Input your command:  
>    -->  
>    ```  
>-   If you use SSH to remotely log in to the Linux or macOS for running obsutil commands, you are advised to set  **TMOUT=0**  to prevent the program from exiting due to the expiration of the SSH session.  

