# Resuming a Failed Upload Task<a name="EN-US_TOPIC_0145591176"></a>

## Function<a name="section1479112110815"></a>

You can use this command to resume a failed upload task based on the task ID.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil cp -recover=xxx [-arcDir=xxx] [-dryRun] [-f] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil cp -recover=xxx [-arcDir=xxx] [-dryRun] [-f] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="15.151515151515152%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="26.262626262626267%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="58.58585858585859%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row382993520717"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p28307351476"><a name="p28307351476"></a><a name="p28307351476"></a>recover</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p198304351672"><a name="p198304351672"></a><a name="p198304351672"></a>Mandatory (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p188308358715"><a name="p188308358715"></a><a name="p188308358715"></a>ID of the upload task to be resumed</p>
<div class="note" id="note10332165610814"><a name="note10332165610814"></a><a name="note10332165610814"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul625425894"></a><a name="ul625425894"></a><ul id="ul625425894"><li>You can obtain the task ID after an upload task is complete, or query it based on the file name of the operation result list, which is the 36 characters excluding the suffix <strong id="b18232828387"><a name="b18232828387"></a><a name="b18232828387"></a>.txt</strong> in the file name.</li><li>You can locate the upload task to be resumed in the directory where the result lists reside. For details about the directory of the result lists, see additional parameter <strong id="b1845116713389"><a name="b1845116713389"></a><a name="b1845116713389"></a>o</strong>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row108730315315"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p7755144410481"><a name="p7755144410481"></a><a name="p7755144410481"></a>arcDir</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p13755134434810"><a name="p13755134434810"></a><a name="p13755134434810"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1175594410487"><a name="p1175594410487"></a><a name="p1175594410487"></a>Path to which the uploaded files are archived</p>
</td>
</tr>
<tr id="row1966219101713"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p156631810914"><a name="p156631810914"></a><a name="p156631810914"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p10663610617"><a name="p10663610617"></a><a name="p10663610617"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p76630101412"><a name="p76630101412"></a><a name="p76630101412"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row20808111443516"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p78083143357"><a name="p78083143357"></a><a name="p78083143357"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p1480891414357"><a name="p1480891414357"></a><a name="p1480891414357"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p20808141443516"><a name="p20808141443516"></a><a name="p20808141443516"></a>Indicates incremental upload. If this parameter is set, each file can be uploaded only when it does not exist in the bucket, its size is different from the namesake one in the bucket, or it has the latest modification time.</p>
</td>
</tr>
<tr id="row14823182511325"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p10823125153219"><a name="p10823125153219"></a><a name="p10823125153219"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p12641397348"><a name="p12641397348"></a><a name="p12641397348"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p582315257326"><a name="p582315257326"></a><a name="p582315257326"></a>After the upload is complete, check whether the sizes of the objects in the bucket are the same as those of the local files.</p>
</td>
</tr>
<tr id="row77547237323"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p16754102323216"><a name="p16754102323216"></a><a name="p16754102323216"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p9286153913341"><a name="p9286153913341"></a><a name="p9286153913341"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p10755202313327"><a name="p10755202313327"></a><a name="p10755202313327"></a>After the upload completes, check whether the MD5 values of the objects in the bucket are the same as those of the local files.</p>
<div class="note" id="note9812411501"><a name="note9812411501"></a><a name="note9812411501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul4628165215267"></a><a name="ul4628165215267"></a><ul id="ul4628165215267"><li>If the size of the file or folder to be uploaded is too large, using this parameter will degrade the overall performance due to MD5 calculation.</li><li>After the MD5 value verification is successful, the parameter value is set to the object metadata <strong id="b040322414717"><a name="b040322414717"></a><a name="b040322414717"></a>x-obs-md5chksum</strong>, which is used for later MD5 verification during download or copy.</li></ul>
</div></div>
</td>
</tr>
<tr id="row14988181814356"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p1098811817356"><a name="p1098811817356"></a><a name="p1098811817356"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p14989181813352"><a name="p14989181813352"></a><a name="p14989181813352"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p49897182353"><a name="p49897182353"></a><a name="p49897182353"></a>Indicates the maximum number of concurrent multipart upload tasks when uploading a file. The default value is the value of <strong id="b19911849151313"><a name="b19911849151313"></a><a name="b19911849151313"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row7456112216356"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p1045612212352"><a name="p1045612212352"></a><a name="p1045612212352"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p74561122123519"><a name="p74561122123519"></a><a name="p74561122123519"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p114561522173511"><a name="p114561522173511"></a><a name="p114561522173511"></a>Indicates the threshold for enabling multipart upload, in bytes. The default value is the value of <strong id="b1139713358465"><a name="b1139713358465"></a><a name="b1139713358465"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note956121318501"><a name="note956121318501"></a><a name="note956121318501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the file or folder to be uploaded is smaller than the threshold, upload it directly. Otherwise, a multipart upload is required.</li><li>If you upload a file or folder directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b1832546105014"><a name="b1832546105014"></a><a name="b1832546105014"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row13242162823512"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p324212285352"><a name="p324212285352"></a><a name="p324212285352"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p82428286353"><a name="p82428286353"></a><a name="p82428286353"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified when uploading files. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding three values indicate private read and write, public read, and public read and write.</p>
</div></div>
</td>
</tr>
<tr id="row10780184745015"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Indicates the storage classes of objects that can be specified when uploading files. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b1769934211230"><a name="b1769934211230"></a><a name="b1769934211230"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b05680571972"><a name="b05680571972"></a><a name="b05680571972"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b06902115818"><a name="b06902115818"></a><a name="b06902115818"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row15476193193517"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p15476113117359"><a name="p15476113117359"></a><a name="p15476113117359"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p3476173116357"><a name="p3476173116357"></a><a name="p3476173116357"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1447673143520"><a name="p1447673143520"></a><a name="p1447673143520"></a>Indicates the customized metadata that can be specified when uploading files. The format is <strong id="b5986113545112"><a name="b5986113545112"></a><a name="b5986113545112"></a>key1:value1#key2:value2#key3:value3</strong>.</p>
<div class="note" id="note2863162085514"><a name="note2863162085514"></a><a name="note2863162085514"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p686342013559"><a name="p686342013559"></a><a name="p686342013559"></a>The preceding value indicates that the object in the bucket contains three groups of customized metadata after the file is uploaded: <strong id="b89423875116"><a name="b89423875116"></a><a name="b89423875116"></a>key1:value1</strong>, <strong id="b5950387516"><a name="b5950387516"></a><a name="b5950387516"></a>key2:value2</strong>, and <strong id="b179773875113"><a name="b179773875113"></a><a name="b179773875113"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row1898318335357"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p139834336354"><a name="p139834336354"></a><a name="p139834336354"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p20983173343511"><a name="p20983173343511"></a><a name="p20983173343511"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p161953211218"><a name="p161953211218"></a><a name="p161953211218"></a>Indicates the size of each part in a multipart upload task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b4994850124714"><a name="b4994850124714"></a><a name="b4994850124714"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note645113418434"><a name="note645113418434"></a><a name="note645113418434"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul177981443114514"></a><a name="ul177981443114514"></a><ul id="ul177981443114514"><li>This value can contain a capacity unit. For example, <strong id="b10124515135210"><a name="b10124515135210"></a><a name="b10124515135210"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b12784171615211"><a name="b12784171615211"></a><a name="b12784171615211"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source file size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row83017445358"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p83084493510"><a name="p83084493510"></a><a name="p83084493510"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p830044143513"><a name="p830044143513"></a><a name="p830044143513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p730944183517"><a name="p730944183517"></a><a name="p730944183517"></a>Indicates the folder where the part records reside. The default value is <strong id="b165914316556"><a name="b165914316556"></a><a name="b165914316556"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note13886557132615"><a name="note13886557132615"></a><a name="note13886557132615"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart upload and saved to the <strong id="b2012163014619"><a name="b2012163014619"></a><a name="b2012163014619"></a>upload</strong> subfolder. After the upload succeeds, its part record is deleted automatically. If the upload fails or is suspended, the system attempts to resume the task according to its part record when you perform the upload the next time.</p>
</div></div>
</td>
</tr>
<tr id="row45081155113516"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p7508105553515"><a name="p7508105553515"></a><a name="p7508105553515"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p202915596214"><a name="p202915596214"></a><a name="p202915596214"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row20127634366"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p51273333618"><a name="p51273333618"></a><a name="p51273333618"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p5361592215"><a name="p5361592215"></a><a name="p5361592215"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p112661531143816"><a name="p112661531143816"></a><a name="p112661531143816"></a>Maximum number of concurrent tasks for uploading a folder. The default value is the value of <strong id="b1837402371316"><a name="b1837402371316"></a><a name="b1837402371316"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row152961440116"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p93091157917"><a name="p93091157917"></a><a name="p93091157917"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p1416597217"><a name="p1416597217"></a><a name="p1416597217"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p63147571114"><a name="p63147571114"></a><a name="p63147571114"></a>Indicates the file matching patterns that are excluded, for example: <strong id="b9216171315555"><a name="b9216171315555"></a><a name="b9216171315555"></a>*.txt</strong>.</p>
<div class="note" id="note860825419155"><a name="note860825419155"></a><a name="note860825419155"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul83178571118"></a><a name="ul83178571118"></a><ul id="ul83178571118"><li>The asterisk (*) represents any characters, and question mark (?) represents only one character. For instance, <strong id="b1537412213553"><a name="b1537412213553"></a><a name="b1537412213553"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b937612229559"><a name="b937612229559"></a><a name="b937612229559"></a>abc</strong> and ends with <strong id="b113775224553"><a name="b113775224553"></a><a name="b113775224553"></a>.txt</strong>.</li><li>You can use <strong id="b22181224105510"><a name="b22181224105510"></a><a name="b22181224105510"></a>\*</strong> to represent <strong id="b11221824175518"><a name="b11221824175518"></a><a name="b11221824175518"></a>*</strong> and <strong id="b222232495511"><a name="b222232495511"></a><a name="b222232495511"></a>\?</strong> to represent <strong id="b1322322435519"><a name="b1322322435519"></a><a name="b1322322435519"></a>?</strong>.</li><li>If the name of the file to be uploaded matches the value of this parameter, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note17161195812153"><a name="note17161195812153"></a><a name="note17161195812153"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern applies only to files in a folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row15404620173616"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p64040209360"><a name="p64040209360"></a><a name="p64040209360"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p24515919220"><a name="p24515919220"></a><a name="p24515919220"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p85831634179"><a name="p85831634179"></a><a name="p85831634179"></a>Indicates the file matching patterns that are included, for example: <strong id="b1523614513578"><a name="b1523614513578"></a><a name="b1523614513578"></a>*.jpg</strong>.</p>
<div class="note" id="note6126191912710"><a name="note6126191912710"></a><a name="note6126191912710"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul144557173539"></a><a name="ul144557173539"></a><ul id="ul144557173539"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b27721812565"><a name="b27721812565"></a><a name="b27721812565"></a>**</strong> to represent <strong id="b127759121062"><a name="b127759121062"></a><a name="b127759121062"></a>*</strong> and <strong id="b12776111211610"><a name="b12776111211610"></a><a name="b12776111211610"></a>\?</strong> to represent <strong id="b97786122613"><a name="b97786122613"></a><a name="b97786122613"></a>?</strong>.</li><li>Only after identifying that the name of the file to be uploaded does not match the value of <strong id="b22581581369"><a name="b22581581369"></a><a name="b22581581369"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is uploaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note7627205718163"><a name="note7627205718163"></a><a name="note7627205718163"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul17172248195614"></a><a name="ul17172248195614"></a><ul id="ul17172248195614"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern applies only to files in a folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1809121510401"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when uploading files. Only files whose last modification time is within the configured time range are uploaded.</p>
<p id="p1879635019517"><a name="p1879635019517"></a><a name="p1879635019517"></a>This pattern has a lower priority than the file matching patterns (<strong id="b37597075317"><a name="b37597075317"></a><a name="b37597075317"></a>exclude</strong>/<strong id="b12760501539"><a name="b12760501539"></a><a name="b12760501539"></a>include</strong>). That is, the time range matching pattern is executed after the configured file matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i1525423112818"><a name="i1525423112818"></a><a name="i1525423112818"></a>time1</em><strong id="b6526323172819"><a name="b6526323172819"></a><a name="b6526323172819"></a>-</strong><em id="i3526192318288"><a name="i3526192318288"></a><a name="i3526192318288"></a>time2</em>, where <em id="i1852711231288"><a name="i1852711231288"></a><a name="i1852711231288"></a>time1</em> must be earlier than or the same as <em id="i1952818234281"><a name="i1952818234281"></a><a name="i1952818234281"></a>time2</em>. The time format is <em id="i1752982319285"><a name="i1752982319285"></a><a name="i1752982319285"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b114382718287"><a name="b114382718287"></a><a name="b114382718287"></a>*-</strong><em id="i17443271283"><a name="i17443271283"></a><a name="i17443271283"></a>time2</em>, all files whose last modification time is earlier than <em id="i245727192819"><a name="i245727192819"></a><a name="i245727192819"></a>time2</em> are matched. If it is set to <em id="i7467273285"><a name="i7467273285"></a><a name="i7467273285"></a>time1</em><strong id="b346427132815"><a name="b346427132815"></a><a name="b346427132815"></a>-*</strong>, all files whose last modification time is later than <em id="i10472279282"><a name="i10472279282"></a><a name="i10472279282"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p95781751171216"><a name="p95781751171216"></a><a name="p95781751171216"></a>Time in the matching pattern is the UTC time.</p>
</div></div>
</td>
</tr>
<tr id="row1610491411911"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b1917013311751"><a name="b1917013311751"></a><a name="b1917013311751"></a>include</strong> or <strong id="b917033115519"><a name="b917033115519"></a><a name="b917033115519"></a>exclude</strong>) and the time matching pattern (<strong id="b517113311754"><a name="b517113311754"></a><a name="b517113311754"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row1015315386301"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p13911642133014"><a name="p13911642133014"></a><a name="p13911642133014"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p84814592220"><a name="p84814592220"></a><a name="p84814592220"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p139151442123019"><a name="p139151442123019"></a><a name="p139151442123019"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b396331195410"><a name="b396331195410"></a><a name="b396331195410"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note114181972212"><a name="note114181972212"></a><a name="note114181972212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b16838421104311"><a name="b16838421104311"></a><a name="b16838421104311"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i18391221134318"><a name="i18391221134318"></a><a name="i18391221134318"></a>time</em><strong id="b983912184313"><a name="b983912184313"></a><a name="b983912184313"></a>_TaskId.txt</strong><p id="p934418291340"><a name="p934418291340"></a><a name="p934418291340"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b93326923919"><a name="b93326923919"></a><a name="b93326923919"></a>recordMaxLogSize</strong> and <strong id="b23331397399"><a name="b23331397399"></a><a name="b23331397399"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row19911919359"><td class="cellrowborder" valign="top" width="15.151515151515152%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="26.262626262626267%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil cp -recover 104786c8-27c2-48fc-bc6a-5886596fb0ed -f**  command to resume the failed upload task.

```
obsutil cp -recover 104786c8-27c2-48fc-bc6a-5886596fb0ed -f

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx

[========================================================] 100.00% 2.02 KB/s 0s
Succeed count is:   5         Failed count is:    0
Metrics [max cost:90 ms, min cost:45 ms, average cost:63.80 ms, average tps:35.71]
Task id is: a628d6da-c562-4a1f-b687-4fa125de0dc3
```

