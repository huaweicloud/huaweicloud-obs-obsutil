# Using obsutil for Resumable Data Transfer<a name="EN-US_TOPIC_0177232036"></a>

obsutil supports resumable data transfer \(upload, download, and copy\) for large files by using the multipart algorithms for upload, download, and copy. You can set the threshold size for starting a multipart upload, download, or copy task based on your actual requirements to resume the upload, download, or copy task if the task fails or is interrupted. You can specify the threshold size for starting a multipart task in either of the following ways:

1.  Specify the  **defaultBigfileThreshold**  parameter in the configuration file.
2.  When running commands for upload, download, copy, incremental synchronization upload, incremental synchronization download, or incremental synchronization copy, you can specify the  **threshold**  parameter at the command level.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   Priority: Command level parameter  **threshold**  has higher priority than the  **defaultBigfileThreshold**  in the configuration file.  
>-   The threshold size of a multipart task applies to single files or objects. When the size of a file or object is greater than the threshold value, the multipart algorithm is applied to the file or object.  
>-   The multipart algorithm and resumable data transfer are forcibly bound together. That is, once the multipart algorithm is used, the resumable data transfer is enabled for the task.  

