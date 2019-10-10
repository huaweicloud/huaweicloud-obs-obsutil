# Performing Initial Configuration<a name="EN-US_TOPIC_0141181357"></a>

Before using obsutil, you need to configure the interconnection between obsutil and OBS, including the endpoint and access keys \(AK and SK\) of OBS. You can use obsutil to perform operations on OBS buckets and objects only after obtaining the OBS authentication.

## Prerequisites<a name="section1544634913019"></a>

-   You have downloaded the obsutil software package, or compiled the source code to generate obsutil. For details, see  [Open-Source Addresses](introduction-to-obsutil.md#section9338140135813).
-   You have obtained the enabled regions and endpoints of OBS. For details, see  [Regions and Endpoints](https://docs.otc.t-systems.com/en-us/endpoint/index.html).
-   You have obtained the access keys \(AK and SK\). For details about how to obtain access keys, see  [Creating Access Keys \(AK and SK\)](creating-access-keys-(ak-and-sk).md).

## Configuration Method<a name="section14490194733410"></a>

 Run the  **config**  command to initialize obsutil. For details about the  **config**  command, see  [Updating a Configuration File](updating-a-configuration-file.md). The following is an example:

-   In Windows

    ```
    obsutil config -interactive
    
    Please input your ak:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your sk:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your endpoint:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your token:
    
    Config file url:
      C:\Users\tools\.obsutilconfig
    
    Update config file successfully!
    ```

-   In Linux OS or macOS

    ```
    ./obsutil config -interactive
    
    Please input your ak:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your sk:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your endpoint:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your token:
    
    Config file url:
      C:\Users\tools\.obsutilconfig
    
    Update config file successfully!
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   After running the preceding commands, a configuration file  **.obsutilconfig**  is automatically generated in the same home directory of the user who executes obsutil commands \(the  **\~**  directory in Linux or macOS, and the  **C:\\Users\\**_<Username\>_  directory in Windows\).  **.obsutilconfig**  contains all the configuration information of obsutil.  
>-   For details about the parameters in the  **.obsutilconfig**  file, see  [Parameter Description](parameter-description.md).  
>-   The  **.obsutilconfig**  file contains the AK and SK information of a user. Therefore, it is hidden by default to prevent key disclosure. To query the file, run the following command in the home directory of the user who executes obsutil commands.  
>    -   In Windows  
>        ```  
>        dir   
>        ```  
>    -   In Linux or macOS  
>        ```  
>        ls -a   
>        ```  
>        or  
>        ```  
>        ls -al  
>        ```  
>-   obsutil encrypts the AK and SK in the  **.obsutilconfig**  file to ensure key security.  

## Checking the Connectivity<a name="section1048170135411"></a>

After the configuration is complete, you can check whether it is correct by running the following commands:

-   In Windows

    ```
    obsutil ls -s
    ```

-   In Linux or macOS

    ```
    ./obsutil ls -s
    ```


Check the configuration result based on the command output:

-   If the command output contains  **Bucket number is:**, the configuration is correct.
-   If the command output contains  **Http status \[403\]**, the access keys are incorrectly configured.
-   If the command output contains  **A connection attempt failed**, then OBS cannot be accessed. In this case, check the network condition.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>If the command output contains  **Http status \[403\]**, you may not have the required permissions for obtaining the bucket list. In this case, further locate the root cause based on the specific situation.  

