# Result Lists<a name="EN-US_TOPIC_0164918059"></a>

## Configuring Result Lists<a name="section16839141103914"></a>

Result lists are generated when batch tasks complete. By default, they are saved to the subfolder  **.obsutil\_output**  in the home directory of the user who executes obsutil commands. You can specify another folder to save them by setting the additional parameter  **-o**  when executing a command.

## Viewing Result Lists<a name="section14566233134617"></a>

Result lists are classified into success, failure, and warning lists. The naming rule is as follows:  _Operation_** \_\{succeed | failed | warning\}\_report\_ **_Time_** \_TaskId.txt**. For example, after a folder is successfully uploaded, the result list is named as follows:  **cp\_succeed\_report\_20190417021908\_fbbc83e3-98ac-4d19-b23a-64023b1e0c34.txt**.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   If the number of successes, failures, or warnings is zero, the corresponding result list is not generated.  
>-   The task ID of a result list is unique for each operation.  
>-   The maximum size of a result list is 30 MB and the maximum number of lists that can be retained is 1024.  

