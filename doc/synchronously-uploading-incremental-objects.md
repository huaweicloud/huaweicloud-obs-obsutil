# Synchronously Uploading Incremental Objects<a name="EN-US_TOPIC_0150706906"></a>

## Function<a name="section68797549353"></a>

This function synchronizes all content in the local source path to the specified target bucket on OBS, ensuring that the content is consistent between the local path and the target bucket. Incremental synchronization has the following meanings: 1\) Increment: Compare the source file with the target object and upload only the source file that has changes. 2\) Synchronization: After the command is executed, ensure that the local source path is a subset of the target bucket specified by OBS. That is, any file in the local source path has its corresponding object in the target bucket on OBS.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>-   Do not change the local file or folder during synchronization. Otherwise, the synchronization may fail or data may be inconsistent.  
>-   Each file can be synchronously uploaded only when it does not exist in the bucket, its size is different from the namesake one in the bucket, or it has the latest modification time.  

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Uploading a file synchronously

        ```
        obsutil sync file_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-vlength] [-vmd5] [-p=1] [-threshold=5248800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-o=xxx] [-cpd=xxx] [-fr] [-config=xxx]
        ```

    -   Uploading a folder synchronously

        ```
        obsutil sync folder_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Uploading a file synchronously

        ```
        ./obsutil sync file_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-vlength] [-vmd5] [-p=1] [-threshold=5248800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-o=xxx] [-cpd=xxx] [-fr] [-config=xxx]
        ```

    -   Uploading a folder synchronously

        ```
        ./obsutil sync folder_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="15.61%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="19.35%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="65.03999999999999%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>file_url</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory for uploading a file synchronously</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Local file path</p>
</td>
</tr>
<tr id="row994520519116"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p1294565714"><a name="p1294565714"></a><a name="p1294565714"></a>folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p2945850113"><a name="p2945850113"></a><a name="p2945850113"></a>Mandatory for uploading a folder synchronously</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p7945859110"><a name="p7945859110"></a><a name="p7945859110"></a>Local folder path</p>
</td>
</tr>
<tr id="row1538192203220"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p154092113211"><a name="p154092113211"></a><a name="p154092113211"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p105401120328"><a name="p105401120328"></a><a name="p105401120328"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p1954062113212"><a name="p1954062113212"></a><a name="p1954062113212"></a>Bucket name</p>
</td>
</tr>
<tr id="row823371133212"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p13233171113215"><a name="p13233171113215"></a><a name="p13233171113215"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p152331511173216"><a name="p152331511173216"></a><a name="p152331511173216"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p16965474526"><a name="p16965474526"></a><a name="p16965474526"></a>Indicates the object name or object name prefix specified when uploading a file synchronously, or the object name prefix specified when uploading a folder synchronously.</p>
<p id="p060018221533"><a name="p060018221533"></a><a name="p060018221533"></a>The rules are as follows:</p>
<a name="ul7190122515538"></a><a name="ul7190122515538"></a><ul id="ul7190122515538"><li>If this parameter is left blank when synchronously uploading a file, the object is uploaded to the root directory of the bucket and the object name is the file name. If the value ends with a slash (/), the value is used as the object name prefix when the file is uploaded, and the object name is the value plus the file name. Otherwise, the file is uploaded with the value as the object name.</li><li>If this parameter is left blank when synchronously uploading a folder, all objects in the root directory of the bucket are the same as the files in the local folder. If this parameter is configured, objects whose name prefix is the configured value are the same as the files in the local folder.</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul1791612918129"></a><a name="ul1791612918129"></a><ul id="ul1791612918129"><li>If the value of this parameter does not end with a slash (/) when synchronously uploading a folder, the obsutil tool automatically adds a slash (/) at the end of the configured value as the object name prefix.</li><li>For details about how to use this parameter, see <a href="synchronous-upload-examples.md">Synchronous Upload Examples</a>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row81423312116"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p123051451513"><a name="p123051451513"></a><a name="p123051451513"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p113071451811"><a name="p113071451811"></a><a name="p113071451811"></a>Optional for synchronously uploading a file (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p73101351816"><a name="p73101351816"></a><a name="p73101351816"></a>Generates an operation result list when synchronously uploading a file.</p>
</td>
</tr>
<tr id="row41295913415"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p7755144410481"><a name="p7755144410481"></a><a name="p7755144410481"></a>arcDir</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p13755134434810"><a name="p13755134434810"></a><a name="p13755134434810"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p1175594410487"><a name="p1175594410487"></a><a name="p1175594410487"></a>Path to which the synchronously uploaded files are archived</p>
</td>
</tr>
<tr id="row20808111443516"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row890711361755"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p990817361553"><a name="p990817361553"></a><a name="p990817361553"></a>link</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p9321174511519"><a name="p9321174511519"></a><a name="p9321174511519"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p890843615516"><a name="p890843615516"></a><a name="p890843615516"></a>Uploads the actual path of the symbolic-link file/folder</p>
<div class="notice" id="note57871443132412"><a name="note57871443132412"></a><a name="note57871443132412"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul19871513171214"></a><a name="ul19871513171214"></a><ul id="ul19871513171214"><li>If this parameter is not specified and the file to be uploaded is a symbolic-link file whose target file does not exist, the exception message "The system cannot find the file specified" will be displayed in Windows OS, while the exception message "No such file or directory" will be displayed in macOS or Linux OS.</li><li>Avoid the symbolic link loop of a folder, otherwise, the upload will exit due to panic. If you do not want the system to panic, set <strong id="b11111014518"><a name="b11111014518"></a><a name="b11111014518"></a>panicForSymbolicLinkCircle</strong> to <strong id="b11211015451"><a name="b11211015451"></a><a name="b11211015451"></a>false</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row14823182511325"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p10823125153219"><a name="p10823125153219"></a><a name="p10823125153219"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p12641397348"><a name="p12641397348"></a><a name="p12641397348"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p582315257326"><a name="p582315257326"></a><a name="p582315257326"></a>After the synchronous upload is complete, check whether the sizes of the objects in the bucket are the same as those of the local files.</p>
</td>
</tr>
<tr id="row77547237323"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p16754102323216"><a name="p16754102323216"></a><a name="p16754102323216"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p9286153913341"><a name="p9286153913341"></a><a name="p9286153913341"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p10755202313327"><a name="p10755202313327"></a><a name="p10755202313327"></a>After the synchronous upload is complete, check whether the MD5 values of the objects in the bucket are the same as those of the local files.</p>
<div class="note" id="note9812411501"><a name="note9812411501"></a><a name="note9812411501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul11953144002612"></a><a name="ul11953144002612"></a><ul id="ul11953144002612"><li>If the size of the file or folder to be uploaded is too large, using this parameter will degrade the overall performance due to MD5 calculation.</li><li>After the MD5 value verification is successful, the parameter value is set to the object metadata <strong id="b152792381645"><a name="b152792381645"></a><a name="b152792381645"></a>x-obs-md5chksum</strong>, which is used for later MD5 verification during download or copy.</li></ul>
</div></div>
</td>
</tr>
<tr id="row14988181814356"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p1098811817356"><a name="p1098811817356"></a><a name="p1098811817356"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p14989181813352"><a name="p14989181813352"></a><a name="p14989181813352"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p49897182353"><a name="p49897182353"></a><a name="p49897182353"></a>Indicates the maximum number of concurrent multipart upload tasks when uploading a file. The default value is the value of <strong id="b11146123118378"><a name="b11146123118378"></a><a name="b11146123118378"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row7456112216356"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p1045612212352"><a name="p1045612212352"></a><a name="p1045612212352"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p74561122123519"><a name="p74561122123519"></a><a name="p74561122123519"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p114561522173511"><a name="p114561522173511"></a><a name="p114561522173511"></a>Indicates the threshold for enabling multipart upload, in bytes. The default value is the value of <strong id="b1139713358465"><a name="b1139713358465"></a><a name="b1139713358465"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note956121318501"><a name="note956121318501"></a><a name="note956121318501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the file or folder to be uploaded is smaller than the threshold, upload it directly. Otherwise, a multipart upload is required.</li><li>If you upload a file or folder directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b119361641103410"><a name="b119361641103410"></a><a name="b119361641103410"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row13242162823512"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p324212285352"><a name="p324212285352"></a><a name="p324212285352"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p82428286353"><a name="p82428286353"></a><a name="p82428286353"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified when synchronously uploading files. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li><li>bucket-owner-full-control</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding four values indicate private read and write, public read, public read and write, and bucket owner full control.</p>
</div></div>
</td>
</tr>
<tr id="row10780184745015"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Indicates the storage classes of objects that can be specified when synchronously uploading files. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b78111078453"><a name="b78111078453"></a><a name="b78111078453"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b1411013201656"><a name="b1411013201656"></a><a name="b1411013201656"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b182208411767"><a name="b182208411767"></a><a name="b182208411767"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row15476193193517"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p15476113117359"><a name="p15476113117359"></a><a name="p15476113117359"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p3476173116357"><a name="p3476173116357"></a><a name="p3476173116357"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p1447673143520"><a name="p1447673143520"></a><a name="p1447673143520"></a>Indicates the customized metadata that can be specified when uploading files. The format is <strong id="b18642117153113"><a name="b18642117153113"></a><a name="b18642117153113"></a>key1:value1#key2:value2#key3:value3</strong>.</p>
<div class="note" id="note2863162085514"><a name="note2863162085514"></a><a name="note2863162085514"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p686342013559"><a name="p686342013559"></a><a name="p686342013559"></a>The preceding value indicates that the object in the bucket contains three groups of customized metadata after the file is uploaded: <strong id="b89423875116"><a name="b89423875116"></a><a name="b89423875116"></a>key1:value1</strong>, <strong id="b5950387516"><a name="b5950387516"></a><a name="b5950387516"></a>key2:value2</strong>, and <strong id="b179773875113"><a name="b179773875113"></a><a name="b179773875113"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row1898318335357"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p139834336354"><a name="p139834336354"></a><a name="p139834336354"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p20983173343511"><a name="p20983173343511"></a><a name="p20983173343511"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p161953211218"><a name="p161953211218"></a><a name="p161953211218"></a>Indicates the size of each part in a multipart upload task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b8805827111319"><a name="b8805827111319"></a><a name="b8805827111319"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note0518193012426"><a name="note0518193012426"></a><a name="note0518193012426"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul835113517435"></a><a name="ul835113517435"></a><ul id="ul835113517435"><li>This value can contain a capacity unit. For example, <strong id="b187791637154615"><a name="b187791637154615"></a><a name="b187791637154615"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b2058883394616"><a name="b2058883394616"></a><a name="b2058883394616"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source file size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row83017445358"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p83084493510"><a name="p83084493510"></a><a name="p83084493510"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p830044143513"><a name="p830044143513"></a><a name="p830044143513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p730944183517"><a name="p730944183517"></a><a name="p730944183517"></a>Indicates the folder where the part records reside. The default value is <strong id="b154521528192415"><a name="b154521528192415"></a><a name="b154521528192415"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note13886557132615"><a name="note13886557132615"></a><a name="note13886557132615"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart upload and saved to the <strong id="b512314618139"><a name="b512314618139"></a><a name="b512314618139"></a>upload</strong> subfolder. After the upload succeeds, its part record is deleted automatically. If the upload fails or is suspended, the system attempts to resume the task according to its part record when you perform the upload the next time.</p>
</div></div>
</td>
</tr>
<tr id="row20127634366"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p51273333618"><a name="p51273333618"></a><a name="p51273333618"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p989613496247"><a name="p989613496247"></a><a name="p989613496247"></a>Optional for synchronously uploading a folder (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p112661531143816"><a name="p112661531143816"></a><a name="p112661531143816"></a>Indicates the maximum number of concurrent tasks for uploading a folder synchronously. The default value is the value of <strong id="b13318131385"><a name="b13318131385"></a><a name="b13318131385"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row152961440116"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p93091157917"><a name="p93091157917"></a><a name="p93091157917"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p1631217571015"><a name="p1631217571015"></a><a name="p1631217571015"></a>Optional for synchronously uploading a folder (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p63147571114"><a name="p63147571114"></a><a name="p63147571114"></a>Indicates the file matching patterns that are excluded, for example: <strong id="b1700257182615"><a name="b1700257182615"></a><a name="b1700257182615"></a>*.txt</strong>.</p>
<div class="note" id="note860825419155"><a name="note860825419155"></a><a name="note860825419155"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul83178571118"></a><a name="ul83178571118"></a><ul id="ul83178571118"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b19146165942616"><a name="b19146165942616"></a><a name="b19146165942616"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b10147759172611"><a name="b10147759172611"></a><a name="b10147759172611"></a>abc</strong> and ends with <strong id="b15148105942616"><a name="b15148105942616"></a><a name="b15148105942616"></a>.txt</strong>.</li><li>You can use <strong id="b240614242715"><a name="b240614242715"></a><a name="b240614242715"></a>\*</strong> to represent <strong id="b44078212718"><a name="b44078212718"></a><a name="b44078212718"></a>*</strong> and <strong id="b140812102715"><a name="b140812102715"></a><a name="b140812102715"></a>\?</strong> to represent <strong id="b1440992132711"><a name="b1440992132711"></a><a name="b1440992132711"></a>?</strong>.</li><li>If the name of the file to be uploaded matches the value of this parameter, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note17161195812153"><a name="note17161195812153"></a><a name="note17161195812153"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern takes effect only for files in the folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row15404620173616"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p64040209360"><a name="p64040209360"></a><a name="p64040209360"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p1856190112516"><a name="p1856190112516"></a><a name="p1856190112516"></a>Optional for synchronously uploading a folder (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p85831634179"><a name="p85831634179"></a><a name="p85831634179"></a>Indicates the file matching patterns that are included, for example: <strong id="b1936152316275"><a name="b1936152316275"></a><a name="b1936152316275"></a>*.jpg</strong>.</p>
<div class="note" id="note6126191912710"><a name="note6126191912710"></a><a name="note6126191912710"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul144557173539"></a><a name="ul144557173539"></a><ul id="ul144557173539"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b17420426142714"><a name="b17420426142714"></a><a name="b17420426142714"></a>\*</strong> to represent <strong id="b8421122632715"><a name="b8421122632715"></a><a name="b8421122632715"></a>*</strong> and <strong id="b16422202622719"><a name="b16422202622719"></a><a name="b16422202622719"></a>\?</strong> to represent <strong id="b7423726122711"><a name="b7423726122711"></a><a name="b7423726122711"></a>?</strong>.</li><li>Only after identifying that the name of the file to be uploaded does not match the value of <strong id="b22581581369"><a name="b22581581369"></a><a name="b22581581369"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is uploaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note7627205718163"><a name="note7627205718163"></a><a name="note7627205718163"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul1024816334581"></a><a name="ul1024816334581"></a><ul id="ul1024816334581"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern takes effect only for files in the folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1258034674412"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional for synchronously uploading a folder (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when synchronously uploading files. Only files whose last modification time is within the configured time range are uploaded.</p>
<p id="p1467441919542"><a name="p1467441919542"></a><a name="p1467441919542"></a>This pattern has a lower priority than the file matching patterns (<strong id="b154809501181"><a name="b154809501181"></a><a name="b154809501181"></a>exclude</strong>/<strong id="b34822501581"><a name="b34822501581"></a><a name="b34822501581"></a>include</strong>). That is, the time range matching pattern is executed after the configured file matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i13561927173114"><a name="i13561927173114"></a><a name="i13561927173114"></a>time1</em><strong id="b0357172733120"><a name="b0357172733120"></a><a name="b0357172733120"></a>-</strong><em id="i03571276318"><a name="i03571276318"></a><a name="i03571276318"></a>time2</em>, where <em id="i2035820277313"><a name="i2035820277313"></a><a name="i2035820277313"></a>time1</em> must be earlier than or the same as <em id="i193593270311"><a name="i193593270311"></a><a name="i193593270311"></a>time2</em>. The time format is <em id="i7359102715317"><a name="i7359102715317"></a><a name="i7359102715317"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b697622914312"><a name="b697622914312"></a><a name="b697622914312"></a>*-</strong><em id="i897611299316"><a name="i897611299316"></a><a name="i897611299316"></a>time2</em>, all files whose last modification time is earlier than <em id="i1097792943117"><a name="i1097792943117"></a><a name="i1097792943117"></a>time2</em> are matched. If it is set to <em id="i199770297314"><a name="i199770297314"></a><a name="i199770297314"></a>time1</em><strong id="b89783292318"><a name="b89783292318"></a><a name="b89783292318"></a>-*</strong>, all files whose last modification time is later than <em id="i1697822963118"><a name="i1697822963118"></a><a name="i1697822963118"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p3426114111312"><a name="p3426114111312"></a><a name="p3426114111312"></a>Time in the matching pattern is the UTC time.</p>
</div></div>
</td>
</tr>
<tr id="row182121926282"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b935410221945"><a name="b935410221945"></a><a name="b935410221945"></a>include</strong> or <strong id="b193551822845"><a name="b193551822845"></a><a name="b193551822845"></a>exclude</strong>) and the time matching pattern (<strong id="b635519221842"><a name="b635519221842"></a><a name="b635519221842"></a>timeRange</strong>) also take effect on folders.</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row1015315386301"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p13911642133014"><a name="p13911642133014"></a><a name="p13911642133014"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p0913194217303"><a name="p0913194217303"></a><a name="p0913194217303"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p139151442123019"><a name="p139151442123019"></a><a name="p139151442123019"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b20803754122412"><a name="b20803754122412"></a><a name="b20803754122412"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note114181972212"><a name="note114181972212"></a><a name="note114181972212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b84361234105415"><a name="b84361234105415"></a><a name="b84361234105415"></a>sync_{succeed  | failed | warning}_report_</strong><em id="i57015165411"><a name="i57015165411"></a><a name="i57015165411"></a>time</em><strong id="b481515205412"><a name="b481515205412"></a><a name="b481515205412"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b9267195911367"><a name="b9267195911367"></a><a name="b9267195911367"></a>recordMaxLogSize</strong> and <strong id="b112681459113611"><a name="b112681459113611"></a><a name="b112681459113611"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row5678162683712"><td class="cellrowborder" valign="top" width="15.61%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="19.35%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65.03999999999999%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil sync d:\\temp\\test.txt obs://bucket-test/key**  command to synchronously upload a file.

```
obsutil sync d:\temp\test.txt obs://bucket-test/key

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx

[====================================================] 100.00% 1.68 MB/s 5s
Upload successfully, 8.46MB, d:\temp\test.txt --> obs://bucket-test/key
```

-   Take the Windows OS as an example. Run the  **obsutil sync d:\\temp obs://bucket-test/temp**  command to synchronously upload a folder.

```
obsutil sync d:\temp obs://bucket-test/temp

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx

[========================================================] 100.00% 2.02 KB/s 0s
Succeed count is:   5         Failed count is:    0
Metrics [max cost:90 ms, min cost:45 ms, average cost:63.80 ms, average tps:35.71]
Task id is: 104786c8-27c2-48fc-bc6a-5886596fb0ed
```

-   For more examples, see  [Synchronous Upload Examples](synchronous-upload-examples.md).

