# Log Files<a name="EN-US_TOPIC_0164918058"></a>

## Configuring Log Files<a name="section16839141103914"></a>

obsutil log files include tool logs and SDK logs. You can add the following parameters to the  **.obsconfigutil**  file to enable the two logging functions.

-   Tool logging \(records the log information generated during obsutil running\): configure  **utilLogPath**,  **utilLogBackups**,  **utilLogLevel**, and  **utilMaxLogSize**.
-   SDK logging \(records the log information generated when using obsutil to call OBS server-side APIs\): configure  **sdkLogPath**,  **sdkLogBackups**,  **sdkLogLevel**, and  **sdkMaxLogSize**.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   For details about the parameter description, see  [Parameter Description](parameter-description.md).  
>-   **utilLogPath**  and  **sdkLogPath**  indicate the absolute paths of the log files, not the folders that store the log files.  
>-   If  **utilLogPath**  and  **sdkLogPath**  are not specified, tool logging and SDK logging are not enabled, and therefore no log file is generated during obsutil running.  
>-   Log files that are rolled over are named as follows:  _filename_**.log.**_number_  

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>If multiple obsutil processes are running at the same time, log files may fail to be written concurrently or may be lost. In this case, add parameter  **-config**  when running commands to configure an independent configuration file for each process. Make sure that  **utilLogPath**  and  **sdkLogPath**  are set to different paths for each process.  

## Collecting Log Files<a name="section53132393016"></a>

You can collect logs in either of the following methods:

Method 1: Use auxiliary commands by referring to  [Archiving Log Files](archiving-log-files.md).

Method 2: Locate the paths specified by  **utilLogPath**  and  **sdkLogPath**  in the configuration file, and then search for the log files in the corresponding paths in the local file system.

