# Uploading an Object<a name="EN-US_TOPIC_0141181368"></a>

## Function<a name="section1479112110815"></a>

You can use this command to upload one or more local files or folders to a specified path in OBS. These files can be texts, images, videos, or any other type of files.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>Do not change the local file or folder when uploading it. Otherwise, the upload may fail or data may be inconsistent.  

## Restrictions<a name="section144331226124017"></a>

obsutil has restrictions on the size of files or folders to be uploaded. You can upload an empty file or folder of 0 bytes. You can also upload a single file or folder with a maximum size of 5 GB in normal mode or a single file or folder with a maximum size of 48.8 TB in multipart mode.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Uploading a file

        ```
        obsutil cp file_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=5248800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-o=xxx] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Uploading a folder

        ```
        obsutil cp folder_url obs://bucket[/key] -r [-arcDir=xxx] [-dryRun] [-link] [-f] [-flat] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```

    -   Uploading multiple files/folders

        ```
        obsutil cp file1_url,folder1_url|filelist_url obs://bucket[/prefix] -msm=1 [-r] [-arcDir=xxx] [-dryRun] [-link] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx][-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Uploading a file

        ```
        ./obsutil cp file_url obs://bucket[/key] [-arcDir=xxx] [-dryRun] [-link] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=5248800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-o=xxx] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Uploading a folder

        ```
        ./obsutil cp folder_url obs://bucket[/key] -r [-arcDir=xxx] [-dryRun] [-link] [-f] [-flat] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```

    -   Uploading multiple files/folders

        ```
        ./obsutil cp file1_url,folder1_url|filelist_url obs://bucket[/prefix] -msm=1 [-r] [-arcDir=xxx] [-dryRun] [-link] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx][-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16.161616161616163%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="25.252525252525253%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="58.58585858585859%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>file_url</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p271644516461"><a name="p271644516461"></a><a name="p271644516461"></a>Optional for uploading multiple files/folders</p>
<p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory for uploading a file</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Local file path</p>
<div class="note" id="note18519241476"><a name="note18519241476"></a><a name="note18519241476"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul3212857124912"></a><a name="ul3212857124912"></a><ul id="ul3212857124912"><li>No paths can be nested when uploading multiple files/folders. For example, you cannot specify <strong id="b1848971413517"><a name="b1848971413517"></a><a name="b1848971413517"></a>/a/b/c</strong> and <strong id="b435701814519"><a name="b435701814519"></a><a name="b435701814519"></a>/a/b/</strong> at the same time.</li><li>If this parameter is configured when uploading multiple files/folders, <strong id="b9663111914354"><a name="b9663111914354"></a><a name="b9663111914354"></a>msm</strong> must be set to <strong id="b79301721123515"><a name="b79301721123515"></a><a name="b79301721123515"></a>1</strong>. In this case, use commas (,) to separate multiple file paths, for example, <strong id="b11255202563616"><a name="b11255202563616"></a><a name="b11255202563616"></a>file_url1,file_url2</strong>.</li><li>Files and folders can both be included when uploading multiple files/folders. For example, <strong id="b16921184147"><a name="b16921184147"></a><a name="b16921184147"></a>file_url1,folder_url1,file_url2,folder_url2</strong>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row994520519116"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1294565714"><a name="p1294565714"></a><a name="p1294565714"></a>folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p15401134114611"><a name="p15401134114611"></a><a name="p15401134114611"></a>Optional for uploading multiple files/folders</p>
<p id="p2945850113"><a name="p2945850113"></a><a name="p2945850113"></a>Mandatory for uploading a folder</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p7945859110"><a name="p7945859110"></a><a name="p7945859110"></a>Local folder path</p>
<div class="note" id="note160711241361"><a name="note160711241361"></a><a name="note160711241361"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul27561826114818"></a><a name="ul27561826114818"></a><ul id="ul27561826114818"><li>If <strong id="b153351113104718"><a name="b153351113104718"></a><a name="b153351113104718"></a>flat</strong> is not configured when uploading a folder, the entire folder is uploaded. If <strong id="b393113319502"><a name="b393113319502"></a><a name="b393113319502"></a>flat</strong> is configured, all files in the folder are uploaded.</li><li>No paths can be nested when uploading multiple files/folders. For example, you cannot specify <strong id="b104891122074"><a name="b104891122074"></a><a name="b104891122074"></a>/a/b/c</strong> and <strong id="b11491721279"><a name="b11491721279"></a><a name="b11491721279"></a>/a/b/</strong> at the same time.</li><li>If this parameter is configured when uploading multiple files/folders, <strong id="b6871116144811"><a name="b6871116144811"></a><a name="b6871116144811"></a>msm</strong> must be set to <strong id="b108713164487"><a name="b108713164487"></a><a name="b108713164487"></a>1</strong>. In this case, use commas (,) to separate multiple folder paths, for example, <strong id="b18872171614815"><a name="b18872171614815"></a><a name="b18872171614815"></a>folder_url1,folder_url2</strong>.</li><li>Files and folders can be included when uploading multiple files/folders. For example, <strong id="b4185156134912"><a name="b4185156134912"></a><a name="b4185156134912"></a>file_url1,folder_url1,file_url2,folder_url2</strong>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row695793184611"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1095713174613"><a name="p1095713174613"></a><a name="p1095713174613"></a>filelist_url</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p11957237464"><a name="p11957237464"></a><a name="p11957237464"></a>Optional for uploading multiple files/folders</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p4957133194611"><a name="p4957133194611"></a><a name="p4957133194611"></a>Indicates the path of the file that contains the list of files/folders to be uploaded. If this parameter is configured, <strong id="b1016475513516"><a name="b1016475513516"></a><a name="b1016475513516"></a>msm</strong> must be set to <strong id="b164781558155"><a name="b164781558155"></a><a name="b164781558155"></a>2</strong>.</p>
<div class="note" id="note482195791219"><a name="note482195791219"></a><a name="note482195791219"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul1653583714502"></a><a name="ul1653583714502"></a><ul id="ul1653583714502"><li>The list file is in common text file formats, such as TXT and CSV. Each line in the file indicates a file or folder to be uploaded. For example:<p id="p15839571121"><a name="p15839571121"></a><a name="p15839571121"></a>file_url1</p>
<p id="p48317576128"><a name="p48317576128"></a><a name="p48317576128"></a>file_url2</p>
<p id="p283205711220"><a name="p283205711220"></a><a name="p283205711220"></a>folder_url1</p>
<p id="p8839574126"><a name="p8839574126"></a><a name="p8839574126"></a>folder_url2</p>
</li><li>No paths can be nested in the list file. For example, you cannot specify <strong id="b09236311376"><a name="b09236311376"></a><a name="b09236311376"></a>/a/b/c</strong> and <strong id="b79238319711"><a name="b79238319711"></a><a name="b79238319711"></a>/a/b/</strong> at the same time.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1538192203220"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p154092113211"><a name="p154092113211"></a><a name="p154092113211"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p105401120328"><a name="p105401120328"></a><a name="p105401120328"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1954062113212"><a name="p1954062113212"></a><a name="p1954062113212"></a>Bucket name</p>
</td>
</tr>
<tr id="row823371133212"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p13233171113215"><a name="p13233171113215"></a><a name="p13233171113215"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p152331511173216"><a name="p152331511173216"></a><a name="p152331511173216"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p16965474526"><a name="p16965474526"></a><a name="p16965474526"></a>Indicates the object name or object name prefix specified when uploading a file, or the object name prefix specified when uploading a folder.</p>
<p id="p060018221533"><a name="p060018221533"></a><a name="p060018221533"></a>The rules are as follows:</p>
<a name="ul7190122515538"></a><a name="ul7190122515538"></a><ul id="ul7190122515538"><li>If this parameter is left blank when uploading a file, the file is uploaded to the root directory of the bucket and the object name is the file name. If the value ends with a slash (/), the value is used as the object name prefix when the file is uploaded, and the object name is the value plus the file name. If the value does not end with a slash (/), the file is uploaded with the value as the object name.</li><li>If this parameter is left blank when uploading a folder, the folder is uploaded to the root directory of the bucket. If the value ends with a slash (/), the value is used as the object name prefix of the folder to be uploaded. If the value does not end with a slash (/), the folder to be uploaded is prefixed with the value plus a slash (/).</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p54049131261"><a name="p54049131261"></a><a name="p54049131261"></a>For details about how to use this parameter, see <a href="upload-examples.md">Upload Examples</a>.</p>
</div></div>
</td>
</tr>
<tr id="row81423312116"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p123051451513"><a name="p123051451513"></a><a name="p123051451513"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p113071451811"><a name="p113071451811"></a><a name="p113071451811"></a>Optional for uploading a file (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p73101351816"><a name="p73101351816"></a><a name="p73101351816"></a>Generates an operation result list when uploading a file.</p>
</td>
</tr>
<tr id="row1251145411111"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p175391756612"><a name="p175391756612"></a><a name="p175391756612"></a>flat</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p11542185610113"><a name="p11542185610113"></a><a name="p11542185610113"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p135464561110"><a name="p135464561110"></a><a name="p135464561110"></a>Uploads all files in a folder but not the folder itself.</p>
</td>
</tr>
<tr id="row1875517442484"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p7755144410481"><a name="p7755144410481"></a><a name="p7755144410481"></a>arcDir</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p13755134434810"><a name="p13755134434810"></a><a name="p13755134434810"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1175594410487"><a name="p1175594410487"></a><a name="p1175594410487"></a>Path to which the uploaded files are archived</p>
</td>
</tr>
<tr id="row1460211592532"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row91758361333"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p8177143613314"><a name="p8177143613314"></a><a name="p8177143613314"></a>link</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1817793616310"><a name="p1817793616310"></a><a name="p1817793616310"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p51771036433"><a name="p51771036433"></a><a name="p51771036433"></a>Uploads the actual path of the symbolic-link file/folder</p>
<div class="notice" id="note54211213239"><a name="note54211213239"></a><a name="note54211213239"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul42762549136"></a><a name="ul42762549136"></a><ul id="ul42762549136"><li>If this parameter is not specified and the file to be uploaded is a symbolic-link file whose target file does not exist, the exception message "The system cannot find the file specified" will be displayed in Windows OS, while the exception message "No such file or directory" will be displayed in macOS or Linux OS.</li><li>Avoid the symbolic link loop of a folder, otherwise, the upload will exit due to panic. If you do not want the system to panic, set <strong id="b1626151403711"><a name="b1626151403711"></a><a name="b1626151403711"></a>panicForSymbolicLinkCircle</strong> to <strong id="b89376169376"><a name="b89376169376"></a><a name="b89376169376"></a>false</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row20808111443516"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p78083143357"><a name="p78083143357"></a><a name="p78083143357"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1480891414357"><a name="p1480891414357"></a><a name="p1480891414357"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p20808141443516"><a name="p20808141443516"></a><a name="p20808141443516"></a>Indicates incremental upload. If this parameter is set, each file can be uploaded only when it does not exist in the bucket, its size is different from the namesake one in the bucket, or it has the latest modification time.</p>
</td>
</tr>
<tr id="row14823182511325"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p10823125153219"><a name="p10823125153219"></a><a name="p10823125153219"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p12641397348"><a name="p12641397348"></a><a name="p12641397348"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p582315257326"><a name="p582315257326"></a><a name="p582315257326"></a>After the upload is complete, check whether the sizes of the objects in the bucket are the same as those of the local files.</p>
</td>
</tr>
<tr id="row77547237323"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p16754102323216"><a name="p16754102323216"></a><a name="p16754102323216"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p9286153913341"><a name="p9286153913341"></a><a name="p9286153913341"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p10755202313327"><a name="p10755202313327"></a><a name="p10755202313327"></a>After the upload completes, check whether the MD5 values of the objects in the bucket are the same as those of the local files.</p>
<div class="note" id="note9812411501"><a name="note9812411501"></a><a name="note9812411501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul2021494682318"></a><a name="ul2021494682318"></a><ul id="ul2021494682318"><li>If the size of the file or folder to be uploaded is too large, using this parameter will degrade the overall performance due to MD5 calculation.</li><li>After the MD5 value verification is successful, the parameter value is set to the object metadata <strong id="b55281138203418"><a name="b55281138203418"></a><a name="b55281138203418"></a>x-obs-md5chksum</strong>, which is used for later MD5 verification during download or copy.</li></ul>
</div></div>
</td>
</tr>
<tr id="row14988181814356"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1098811817356"><a name="p1098811817356"></a><a name="p1098811817356"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p14989181813352"><a name="p14989181813352"></a><a name="p14989181813352"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p49897182353"><a name="p49897182353"></a><a name="p49897182353"></a>Indicates the maximum number of concurrent multipart upload tasks when uploading a file. The default value is the value of <strong id="b11172033171820"><a name="b11172033171820"></a><a name="b11172033171820"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row7456112216356"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1045612212352"><a name="p1045612212352"></a><a name="p1045612212352"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p74561122123519"><a name="p74561122123519"></a><a name="p74561122123519"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p114561522173511"><a name="p114561522173511"></a><a name="p114561522173511"></a>Indicates the threshold for enabling multipart upload, in bytes. The default value is the value of <strong id="b12775287240"><a name="b12775287240"></a><a name="b12775287240"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note956121318501"><a name="note956121318501"></a><a name="note956121318501"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the file or folder to be uploaded is smaller than the threshold, upload it directly. Otherwise, a multipart upload is required.</li><li>If you upload a file or folder directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b109921548143218"><a name="b109921548143218"></a><a name="b109921548143218"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row13242162823512"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p324212285352"><a name="p324212285352"></a><a name="p324212285352"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p82428286353"><a name="p82428286353"></a><a name="p82428286353"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified when uploading files. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li><li>bucket-owner-full-control</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding four values indicate private read and write, public read, public read and write, and bucket owner full control.</p>
</div></div>
</td>
</tr>
<tr id="row10780184745015"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Indicates the storage classes of objects that can be specified when uploading files. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b390163043310"><a name="b390163043310"></a><a name="b390163043310"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b94839151339"><a name="b94839151339"></a><a name="b94839151339"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b18878195311"><a name="b18878195311"></a><a name="b18878195311"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row15476193193517"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p15476113117359"><a name="p15476113117359"></a><a name="p15476113117359"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p3476173116357"><a name="p3476173116357"></a><a name="p3476173116357"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1447673143520"><a name="p1447673143520"></a><a name="p1447673143520"></a>Indicates the customized metadata that can be specified when uploading files. The format is <strong id="b073820107443"><a name="b073820107443"></a><a name="b073820107443"></a>key1:value1#key2:value2#key3:value3</strong>.</p>
<div class="note" id="note2863162085514"><a name="note2863162085514"></a><a name="note2863162085514"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1946616159018"><a name="p1946616159018"></a><a name="p1946616159018"></a>The preceding value indicates that the objects in the bucket contain three groups of customized metadata after the file is uploaded: <strong id="b1822379544"><a name="b1822379544"></a><a name="b1822379544"></a>key1:value1</strong>, <strong id="b1213141175417"><a name="b1213141175417"></a><a name="b1213141175417"></a>key2:value2</strong>, and <strong id="b148241614165414"><a name="b148241614165414"></a><a name="b148241614165414"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row1898318335357"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p139834336354"><a name="p139834336354"></a><a name="p139834336354"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p20983173343511"><a name="p20983173343511"></a><a name="p20983173343511"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p161953211218"><a name="p161953211218"></a><a name="p161953211218"></a>Indicates the size of each part in a multipart upload task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b1638015267568"><a name="b1638015267568"></a><a name="b1638015267568"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note0518193012426"><a name="note0518193012426"></a><a name="note0518193012426"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul7782171533813"></a><a name="ul7782171533813"></a><ul id="ul7782171533813"><li>This value can contain a capacity unit. For example, <strong id="b15121036346"><a name="b15121036346"></a><a name="b15121036346"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b10224142385613"><a name="b10224142385613"></a><a name="b10224142385613"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source file size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row83017445358"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p83084493510"><a name="p83084493510"></a><a name="p83084493510"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p830044143513"><a name="p830044143513"></a><a name="p830044143513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p730944183517"><a name="p730944183517"></a><a name="p730944183517"></a>Indicates the folder where the part records reside. The default value is <strong id="b6834171804617"><a name="b6834171804617"></a><a name="b6834171804617"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note13886557132615"><a name="note13886557132615"></a><a name="note13886557132615"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p17679143314315"><a name="p17679143314315"></a><a name="p17679143314315"></a>A part record is generated during a multipart upload and saved to the <strong id="b44770145465"><a name="b44770145465"></a><a name="b44770145465"></a>upload</strong> subfolder. After the upload succeeds, its part record is deleted automatically. If the upload fails or is suspended, the system attempts to resume the task according to its part record when you perform the upload the next time.</p>
</div></div>
</td>
</tr>
<tr id="row1747449123515"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p1047849113513"><a name="p1047849113513"></a><a name="p1047849113513"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1647749153515"><a name="p1647749153515"></a><a name="p1647749153515"></a>Mandatory for uploading a folder (additional parameter)</p>
<p id="p75945581749"><a name="p75945581749"></a><a name="p75945581749"></a>Optional for uploading multiple files/folders</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p119531818122311"><a name="p119531818122311"></a><a name="p119531818122311"></a>Indicates files and subfolders within the folder when uploading a folder recursively.</p>
</td>
</tr>
<tr id="row45081155113516"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p7508105553515"><a name="p7508105553515"></a><a name="p7508105553515"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p450895519351"><a name="p450895519351"></a><a name="p450895519351"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row20127634366"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p51273333618"><a name="p51273333618"></a><a name="p51273333618"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p989613496247"><a name="p989613496247"></a><a name="p989613496247"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p112661531143816"><a name="p112661531143816"></a><a name="p112661531143816"></a>Indicates the maximum number of concurrent tasks for uploading a folder. The default value is the value of <strong id="b98238381222"><a name="b98238381222"></a><a name="b98238381222"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row191914363329"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p119283683216"><a name="p119283683216"></a><a name="p119283683216"></a>msm</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p10192183618322"><a name="p10192183618322"></a><a name="p10192183618322"></a>Mandatory for uploading multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p1023319845920"><a name="p1023319845920"></a><a name="p1023319845920"></a>Enables the mode for uploading multiple files/folders. Possible values are <strong id="b08451816451"><a name="b08451816451"></a><a name="b08451816451"></a>1</strong> and <strong id="b14515134114513"><a name="b14515134114513"></a><a name="b14515134114513"></a>2</strong>.</p>
<div class="note" id="note67632013115912"><a name="note67632013115912"></a><a name="note67632013115912"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul9256520175917"></a><a name="ul9256520175917"></a><ul id="ul9256520175917"><li>If <strong id="b540314016415"><a name="b540314016415"></a><a name="b540314016415"></a>msm</strong> is set to <strong id="b1776017323516"><a name="b1776017323516"></a><a name="b1776017323516"></a>1</strong>, the source URL indicates a list of file/folder names separated by commas.</li><li>If <strong id="b104011690425"><a name="b104011690425"></a><a name="b104011690425"></a>msm</strong> is set to <strong id="b958823585112"><a name="b958823585112"></a><a name="b958823585112"></a>2</strong>, the source URL indicates a file containing a list of file/folder names.</li><li>If the file or folder name already contains commas (,), do not set <strong id="b1118141655517"><a name="b1118141655517"></a><a name="b1118141655517"></a>msm</strong> to <strong id="b73433199551"><a name="b73433199551"></a><a name="b73433199551"></a>1</strong>.</li><li>If parameter <strong id="b1157144611561"><a name="b1157144611561"></a><a name="b1157144611561"></a>r</strong> is not set, the folders in the list will not be uploaded.</li></ul>
</div></div>
</td>
</tr>
<tr id="row152961440116"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p93091157917"><a name="p93091157917"></a><a name="p93091157917"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1631217571015"><a name="p1631217571015"></a><a name="p1631217571015"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p63147571114"><a name="p63147571114"></a><a name="p63147571114"></a>Indicates the file matching patterns that are excluded, for example: <strong id="b9216171315555"><a name="b9216171315555"></a><a name="b9216171315555"></a>*.txt</strong>.</p>
<div class="note" id="note860825419155"><a name="note860825419155"></a><a name="note860825419155"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul83178571118"></a><a name="ul83178571118"></a><ul id="ul83178571118"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b74276145181"><a name="b74276145181"></a><a name="b74276145181"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b165616361198"><a name="b165616361198"></a><a name="b165616361198"></a>abc</strong> and ends with <strong id="b2067555071910"><a name="b2067555071910"></a><a name="b2067555071910"></a>.txt</strong>.</li><li>You can use <strong id="b1122264172018"><a name="b1122264172018"></a><a name="b1122264172018"></a>\*</strong> to represent <strong id="b1438418713214"><a name="b1438418713214"></a><a name="b1438418713214"></a>*</strong> and <strong id="b12516520202117"><a name="b12516520202117"></a><a name="b12516520202117"></a>\?</strong> to represent <strong id="b1333463219216"><a name="b1333463219216"></a><a name="b1333463219216"></a>?</strong>.</li><li>If the name of the file to be uploaded matches the value of this parameter, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note17161195812153"><a name="note17161195812153"></a><a name="note17161195812153"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern takes effect only for files in the folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row15404620173616"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p64040209360"><a name="p64040209360"></a><a name="p64040209360"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1856190112516"><a name="p1856190112516"></a><a name="p1856190112516"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p85831634179"><a name="p85831634179"></a><a name="p85831634179"></a>Indicates the file matching patterns that are included, for example: <strong id="b2035419392404"><a name="b2035419392404"></a><a name="b2035419392404"></a>*.jpg</strong>.</p>
<div class="note" id="note6126191912710"><a name="note6126191912710"></a><a name="note6126191912710"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul144557173539"></a><a name="ul144557173539"></a><ul id="ul144557173539"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b104111547104010"><a name="b104111547104010"></a><a name="b104111547104010"></a>\*</strong> to represent <strong id="b13413447174016"><a name="b13413447174016"></a><a name="b13413447174016"></a>*</strong> and <strong id="b441434710401"><a name="b441434710401"></a><a name="b441434710401"></a>\?</strong> to represent <strong id="b10415647194012"><a name="b10415647194012"></a><a name="b10415647194012"></a>?</strong>.</li><li>Only after identifying that the name of the file to be uploaded does not match the value of <strong id="b367214063712"><a name="b367214063712"></a><a name="b367214063712"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is uploaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note7627205718163"><a name="note7627205718163"></a><a name="note7627205718163"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul11945840155716"></a><a name="ul11945840155716"></a><ul id="ul11945840155716"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute file path (including the file name and file directory).</li><li>The matching pattern takes effect only for files in the folder.</li></ul>
</div></div>
</td>
</tr>
<tr id="row147961525193919"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p10475132743915"><a name="p10475132743915"></a><a name="p10475132743915"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional for uploading a folder or multiple files/folders (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p447514276391"><a name="p447514276391"></a><a name="p447514276391"></a>Indicates the time range matching pattern when uploading files. Only files whose last modification time is within the configured time range are uploaded.</p>
<p id="p138074294535"><a name="p138074294535"></a><a name="p138074294535"></a>This pattern has a lower priority than the file matching patterns (<strong id="b162911291513"><a name="b162911291513"></a><a name="b162911291513"></a>exclude</strong>/<strong id="b1845513312516"><a name="b1845513312516"></a><a name="b1845513312516"></a>include</strong>). That is, the time range matching pattern is executed after the configured file matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i87617512146"><a name="i87617512146"></a><a name="i87617512146"></a>time1</em><strong id="b15762185112149"><a name="b15762185112149"></a><a name="b15762185112149"></a>-</strong><em id="i1976285141414"><a name="i1976285141414"></a><a name="i1976285141414"></a>time2</em>, where <em id="i67630513145"><a name="i67630513145"></a><a name="i67630513145"></a>time1</em> must be earlier than or the same as <em id="i976312512141"><a name="i976312512141"></a><a name="i976312512141"></a>time2</em>. The time format is <em id="i976465113149"><a name="i976465113149"></a><a name="i976465113149"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b24991414181920"><a name="b24991414181920"></a><a name="b24991414181920"></a>*-</strong><em id="i175001214191919"><a name="i175001214191919"></a><a name="i175001214191919"></a>time2</em>, all files whose last modification time is earlier than <em id="i11501914191917"><a name="i11501914191917"></a><a name="i11501914191917"></a>time2</em> are matched. If it is set to <em id="i750119146199"><a name="i750119146199"></a><a name="i750119146199"></a>time1</em><strong id="b850212145197"><a name="b850212145197"></a><a name="b850212145197"></a>-*</strong>, all files whose last modification time is later than <em id="i5502151410192"><a name="i5502151410192"></a><a name="i5502151410192"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p9611192591218"><a name="p9611192591218"></a><a name="p9611192591218"></a>Time in the matching pattern is the UTC time.</p>
</div></div>
</td>
</tr>
<tr id="row81211271351"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b175395333510"><a name="b175395333510"></a><a name="b175395333510"></a>include</strong> or <strong id="b15150138193510"><a name="b15150138193510"></a><a name="b15150138193510"></a>exclude</strong>) and the time matching pattern (<strong id="b181069271354"><a name="b181069271354"></a><a name="b181069271354"></a>timeRange</strong>) also take effect on folders.</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row1015315386301"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p13911642133014"><a name="p13911642133014"></a><a name="p13911642133014"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p0913194217303"><a name="p0913194217303"></a><a name="p0913194217303"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p139151442123019"><a name="p139151442123019"></a><a name="p139151442123019"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b1427311395462"><a name="b1427311395462"></a><a name="b1427311395462"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note114181972212"><a name="note114181972212"></a><a name="note114181972212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b6970184913217"><a name="b6970184913217"></a><a name="b6970184913217"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i1097110491528"><a name="i1097110491528"></a><a name="i1097110491528"></a>time</em><strong id="b189721649923"><a name="b189721649923"></a><a name="b189721649923"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b231552991516"><a name="b231552991516"></a><a name="b231552991516"></a>recordMaxLogSize</strong> and <strong id="b1458123317155"><a name="b1458123317155"></a><a name="b1458123317155"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row979195743411"><td class="cellrowborder" valign="top" width="16.161616161616163%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.58585858585859%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

<a name="table992610203244"></a>
<table><thead align="left"><tr id="row892913208248"><th class="cellrowborder" valign="top" width="21.73%" id="mcps1.1.3.1.1"><p id="p1392992019245"><a name="p1392992019245"></a><a name="p1392992019245"></a>Field</p>
</th>
<th class="cellrowborder" valign="top" width="78.27%" id="mcps1.1.3.1.2"><p id="p19318207246"><a name="p19318207246"></a><a name="p19318207246"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row79320208248"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p1093217203246"><a name="p1093217203246"></a><a name="p1093217203246"></a>Parallel</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p18753151115409"><a name="p18753151115409"></a><a name="p18753151115409"></a>Parameter <strong id="b19168162103414"><a name="b19168162103414"></a><a name="b19168162103414"></a>-p</strong> in the request</p>
</td>
</tr>
<tr id="row1753615812247"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p15536155820245"><a name="p15536155820245"></a><a name="p15536155820245"></a>Jobs</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p575213111407"><a name="p575213111407"></a><a name="p575213111407"></a>Parameter <strong id="b153685413416"><a name="b153685413416"></a><a name="b153685413416"></a>-j</strong> in the request</p>
</td>
</tr>
<tr id="row98643152253"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p8864415102517"><a name="p8864415102517"></a><a name="p8864415102517"></a>Threshold</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p0751611114014"><a name="p0751611114014"></a><a name="p0751611114014"></a>Parameter <strong id="b179171568359"><a name="b179171568359"></a><a name="b179171568359"></a>-threshold</strong> in the request</p>
</td>
</tr>
<tr id="row14992172352513"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p69926234252"><a name="p69926234252"></a><a name="p69926234252"></a>PartSize</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p3250135316288"><a name="p3250135316288"></a><a name="p3250135316288"></a>Parameter <strong id="b995901313915"><a name="b995901313915"></a><a name="b995901313915"></a>-ps</strong> in the request</p>
</td>
</tr>
<tr id="row1272082862513"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p4720528142516"><a name="p4720528142516"></a><a name="p4720528142516"></a>Exclude</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p1874911115406"><a name="p1874911115406"></a><a name="p1874911115406"></a>Parameter <strong id="b17340175311350"><a name="b17340175311350"></a><a name="b17340175311350"></a>-exclude</strong> in the request</p>
</td>
</tr>
<tr id="row48453315255"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p14845531172510"><a name="p14845531172510"></a><a name="p14845531172510"></a>Include</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p87261611104018"><a name="p87261611104018"></a><a name="p87261611104018"></a>Parameter <strong id="b89716773613"><a name="b89716773613"></a><a name="b89716773613"></a>-include</strong> in the request</p>
</td>
</tr>
<tr id="row34663591518"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p164668597512"><a name="p164668597512"></a><a name="p164668597512"></a>TimeRange</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p846715915511"><a name="p846715915511"></a><a name="p846715915511"></a>Parameter <strong id="b917734933"><a name="b917734933"></a><a name="b917734933"></a>-timeRange</strong> in the request</p>
</td>
</tr>
<tr id="row140161714377"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p11402017193717"><a name="p11402017193717"></a><a name="p11402017193717"></a>VerifyLength</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p1440201715372"><a name="p1440201715372"></a><a name="p1440201715372"></a>Parameter <strong id="b14826192963717"><a name="b14826192963717"></a><a name="b14826192963717"></a>-vlength</strong> in the request</p>
</td>
</tr>
<tr id="row582401713379"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p88242017193718"><a name="p88242017193718"></a><a name="p88242017193718"></a>VerifyMd5</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p1782461773716"><a name="p1782461773716"></a><a name="p1782461773716"></a>Parameter <strong id="b197032038173712"><a name="b197032038173712"></a><a name="b197032038173712"></a>-vmd5</strong> in the request</p>
</td>
</tr>
<tr id="row876382643711"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p147636267379"><a name="p147636267379"></a><a name="p147636267379"></a>CheckpointDir</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p6763142603718"><a name="p6763142603718"></a><a name="p6763142603718"></a>Parameter <strong id="b74971747123719"><a name="b74971747123719"></a><a name="b74971747123719"></a>-cpd</strong> in the request</p>
</td>
</tr>
<tr id="row55591628103716"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p1555915282373"><a name="p1555915282373"></a><a name="p1555915282373"></a>OutputDir</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p9559728133714"><a name="p9559728133714"></a><a name="p9559728133714"></a>Parameter <strong id="b3350195715375"><a name="b3350195715375"></a><a name="b3350195715375"></a>-o</strong> in the request</p>
</td>
</tr>
<tr id="row195514813612"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p5956748113617"><a name="p5956748113617"></a><a name="p5956748113617"></a>ArcDir</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p1095674853610"><a name="p1095674853610"></a><a name="p1095674853610"></a>Parameter <strong id="b604549062"><a name="b604549062"></a><a name="b604549062"></a>-arcDir</strong> in the request</p>
</td>
</tr>
<tr id="row1016394416383"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p17163844123820"><a name="p17163844123820"></a><a name="p17163844123820"></a>Succeed count</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p2016364413383"><a name="p2016364413383"></a><a name="p2016364413383"></a>Number of successful tasks</p>
</td>
</tr>
<tr id="row6534144620386"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p1253444614387"><a name="p1253444614387"></a><a name="p1253444614387"></a>Failed count</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p11534114619383"><a name="p11534114619383"></a><a name="p11534114619383"></a>Number of failed tasks</p>
</td>
</tr>
<tr id="row677892931514"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p677822951511"><a name="p677822951511"></a><a name="p677822951511"></a>Skip count</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p668013391033"><a name="p668013391033"></a><a name="p668013391033"></a>Number of tasks that are skipped during incremental upload, download, or copy, and synchronous upload, download, or copy.</p>
<div class="note" id="note5200341631"><a name="note5200341631"></a><a name="note5200341631"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p120211411639"><a name="p120211411639"></a><a name="p120211411639"></a>Skipped tasks are recorded into successful tasks.</p>
</div></div>
</td>
</tr>
<tr id="row12422201513596"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p1642251535915"><a name="p1642251535915"></a><a name="p1642251535915"></a>Warning count</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p342201525916"><a name="p342201525916"></a><a name="p342201525916"></a>Number of tasks that are executed successfully but contain warnings.</p>
<div class="note" id="note1584392192611"><a name="note1584392192611"></a><a name="note1584392192611"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul258175613303"></a><a name="ul258175613303"></a><ul id="ul258175613303"><li>The task for which a warning is generated may be a failure or a success, which needs to be further determined according to the corresponding result list.</li><li>The number of tasks that generate warnings is independent of the number of successful or failed tasks. The total number of tasks is the number of successful tasks plus the number of failed tasks.</li></ul>
</div></div>
</td>
</tr>
<tr id="row849215496713"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p18493144912719"><a name="p18493144912719"></a><a name="p18493144912719"></a>Succeed bytes</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p949320491272"><a name="p949320491272"></a><a name="p949320491272"></a>Number of bytes that are successfully uploaded or downloaded.</p>
</td>
</tr>
<tr id="row12341111614385"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p1734191619388"><a name="p1734191619388"></a><a name="p1734191619388"></a>max cost</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p18341201623819"><a name="p18341201623819"></a><a name="p18341201623819"></a>Maximum duration of all tasks, in ms</p>
</td>
</tr>
<tr id="row19779133153819"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p12779113112387"><a name="p12779113112387"></a><a name="p12779113112387"></a>min cost</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p187791531203811"><a name="p187791531203811"></a><a name="p187791531203811"></a>Minimum duration of all tasks, in ms</p>
</td>
</tr>
<tr id="row205051240163810"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p550514083813"><a name="p550514083813"></a><a name="p550514083813"></a>average cost</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p6505204083811"><a name="p6505204083811"></a><a name="p6505204083811"></a>Average duration of all tasks, in ms</p>
</td>
</tr>
<tr id="row1429016433910"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p142905433917"><a name="p142905433917"></a><a name="p142905433917"></a>average tps</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p3290104113915"><a name="p3290104113915"></a><a name="p3290104113915"></a>The average number of tasks completed per second</p>
</td>
</tr>
<tr id="row537634713394"><td class="cellrowborder" valign="top" width="21.73%" headers="mcps1.1.3.1.1 "><p id="p203761947153916"><a name="p203761947153916"></a><a name="p203761947153916"></a>Task id</p>
</td>
<td class="cellrowborder" valign="top" width="78.27%" headers="mcps1.1.3.1.2 "><p id="p3376144723919"><a name="p3376144723919"></a><a name="p3376144723919"></a>Unique ID of an operation, which is used to search for the result list generated in a batch task</p>
</td>
</tr>
</tbody>
</table>

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil cp d:\\temp\\test.txt obs://bucket-test/key**  command to upload the  **test.txt**  file in the  **temp**  directory in the  **D:**  drive to bucket  **bucket-test**  and rename the file as  **key**.

```
obsutil cp d:\temp\test.txt obs://bucket-test/key

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx

[====================================================] 100.00% 1.68 MB/s 5s
Upload successfully, 8.46MB, d:\temp\test.txt --> obs://bucket-test/key
```

-   Take the Windows OS as an example. Run the  **obsutil cp d:\\temp obs://bucket-test -f -r**  command to recursively upload all files and subfolders in the  **temp**  directory in the  **D:**  drive to the  **temp**  folder in bucket  **bucket-test**.

```
obsutil cp d:\temp obs://bucket-test -f -r

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

-   For more examples, see  [Upload Examples](upload-examples.md).

