# Fine-Tuning obsutil Performance<a name="EN-US_TOPIC_0162143955"></a>

By default, obsutil uploads, downloads, and copies files or objects whose size is greater than 50 MB in multiple parts.  [Table 1](#table8301721164611)  details related parameters in the  **.obsutilconfig**  file.

**Table  1**  Multipart-related parameters

<a name="table8301721164611"></a>
<table><thead align="left"><tr id="row1930332134612"><th class="cellrowborder" valign="top" width="26.51%" id="mcps1.2.3.1.1"><p id="p203035211465"><a name="p203035211465"></a><a name="p203035211465"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="73.49%" id="mcps1.2.3.1.2"><p id="p3303721134611"><a name="p3303721134611"></a><a name="p3303721134611"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row16303112119462"><td class="cellrowborder" valign="top" width="26.51%" headers="mcps1.2.3.1.1 "><p id="p1050292773816"><a name="p1050292773816"></a><a name="p1050292773816"></a>defaultBigfileThreshold</p>
</td>
<td class="cellrowborder" valign="top" width="73.49%" headers="mcps1.2.3.1.2 "><p id="p14525135010317"><a name="p14525135010317"></a><a name="p14525135010317"></a>Indicates the threshold for triggering multipart tasks, in bytes. If the size of a file to be uploaded, downloaded, or copied is greater than the threshold, the file is uploaded, downloaded, or copied in multiple parts. The default value is 50 MB.</p>
</td>
</tr>
<tr id="row19303122184613"><td class="cellrowborder" valign="top" width="26.51%" headers="mcps1.2.3.1.1 "><p id="p077914845118"><a name="p077914845118"></a><a name="p077914845118"></a>defaultPartSize</p>
</td>
<td class="cellrowborder" valign="top" width="73.49%" headers="mcps1.2.3.1.2 "><p id="p45251650113113"><a name="p45251650113113"></a><a name="p45251650113113"></a>Size of each part, in bytes. The default value is <strong id="b196723543534"><a name="b196723543534"></a><a name="b196723543534"></a>auto</strong>.</p>
<div class="note" id="note12525105093118"><a name="note12525105093118"></a><a name="note12525105093118"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul13525175063117"></a><a name="ul13525175063117"></a><ul id="ul13525175063117"><li>For multipart upload and copy, the value ranges from 100 KB to 5 GB.</li><li>For multipart download, the value is unrestricted.</li></ul>
</div></div>
</td>
</tr>
<tr id="row630320214468"><td class="cellrowborder" valign="top" width="26.51%" headers="mcps1.2.3.1.1 "><p id="p4866163316380"><a name="p4866163316380"></a><a name="p4866163316380"></a>defaultParallels</p>
</td>
<td class="cellrowborder" valign="top" width="73.49%" headers="mcps1.2.3.1.2 "><p id="p35261450123120"><a name="p35261450123120"></a><a name="p35261450123120"></a>Maximum number of concurrent tasks in the multipart mode. The default value is 5.</p>
</td>
</tr>
</tbody>
</table>

Generally, multipart tasks not only speed up transmission but also allow you to resume failed tasks. By default, the part size of a multipart task can be automatically adjusted by the obsutil in the  **auto**  mode. In practice, however, to further improve the upload and download performance, you can adjust the part size according to the file size and the network conditions, to obtain the maximum transmission efficiency and ensure the successful completion of a transmission task.

Adjust the number of concurrent tasks in the multipart mode according to the following formula:

**defaultParallels = Min\(Number of CPUs x 2, Object size/defaultPartSize x 1.5\)**

In the upload, download, and copy commands, parameters  **-p**  and  **-ps**  are used to modify the number of concurrent tasks in the multipart mode and part size respectively, and then deliver the multipart task based on the parameter values configured in the command. The default values in the configuration file are used if you do not set them in a command.

Adjust the number of concurrent tasks in the multipart mode according to the following formula:

**p = Min\(Number of CPUs x 2, Object size/ps x 1.5\)**

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   Resources of a running host are limited. Therefore, if the number of concurrent tasks in the multipart mode is set too large, the performance of obsutil upload, download, or copy may deteriorate due to resource switchover and preemption between threads. In this case, you need to adjust the values of  **defaultParallels**  \(**-p**\) and  **defaultPartSize**  \(**-ps**\) based on the actual file size and network status. To perform a pressure test, lower the two values at first, and then gradually increase them to determine the optimal values.  
>-   If the values of  **defaultParallels**  \(**-p**\) and  **defaultPartSize**  \(**-ps**\) are too large, an EOF error may occur due to network instability. In this case, set the two parameters to smaller values.  
>-   If a batch operation is performed, the destination object size can be set to the average size of the objects to be operated.  

