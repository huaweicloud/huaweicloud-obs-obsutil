# Downloading an Object<a name="EN-US_TOPIC_0142358870"></a>

## Function<a name="section1479112110815"></a>

You can use this command to download an object or download objects in batches by object name prefix to your local PC.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>-   Do not change the source objects in the OBS bucket when downloading a single object or objects in batches. Otherwise, the download may fail or data may be inconsistent.  
>-   If the storage class of the object to be copied is  **cold**, you must restore the object to be downloaded first. Otherwise, the download fails.  

## Command Line Structure<a name="section19587175622519"></a>

-   In Windows
    -   Downloading a single object

        ```
        obsutil cp obs://bucket/key file_or_folder_url [-tempFileDir=xxx] [-dryRun] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-ps=auto] [-cpd=xxx][-fr] [-o=xxx] [-config=xxx]
        ```

    -   Downloading objects in batches

        ```
        obsutil cp obs://bucket[/key] folder_url -r [-tempFileDir=xxx] [-dryRun] [-f] [-flat] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Downloading a single object

        ```
        ./obsutil cp obs://bucket/key file_or_folder_url [-tempFileDir=xxx] [-dryRun] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Downloading objects in batches

        ```
        ./obsutil cp obs://bucket[/key] folder_url -r [-tempFileDir=xxx] [-dryRun] [-f] [-flat] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```



## Parameter Description<a name="section175878561255"></a>

<a name="table1155454517538"></a>
<table><thead align="left"><tr id="row6558154516531"><th class="cellrowborder" valign="top" width="17.17171717171717%" id="mcps1.1.4.1.1"><p id="p1656117451534"><a name="p1656117451534"></a><a name="p1656117451534"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="26.26262626262626%" id="mcps1.1.4.1.2"><p id="p155651545195310"><a name="p155651545195310"></a><a name="p155651545195310"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="56.56565656565656%" id="mcps1.1.4.1.3"><p id="p165678456530"><a name="p165678456530"></a><a name="p165678456530"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row756944516535"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p85701245145318"><a name="p85701245145318"></a><a name="p85701245145318"></a>file_or_folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p457314555313"><a name="p457314555313"></a><a name="p457314555313"></a>Mandatory for downloading an object</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p4573114555313"><a name="p4573114555313"></a><a name="p4573114555313"></a>Local file/folder path</p>
</td>
</tr>
<tr id="row195751145115315"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p1576945115318"><a name="p1576945115318"></a><a name="p1576945115318"></a>folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p75772453537"><a name="p75772453537"></a><a name="p75772453537"></a>Mandatory for downloading objects in batches</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p12579194519531"><a name="p12579194519531"></a><a name="p12579194519531"></a>Local folder path</p>
</td>
</tr>
<tr id="row1157912451534"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p058211459536"><a name="p058211459536"></a><a name="p058211459536"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p5582164575310"><a name="p5582164575310"></a><a name="p5582164575310"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p158414515316"><a name="p158414515316"></a><a name="p158414515316"></a>Bucket name</p>
</td>
</tr>
<tr id="row1584114513535"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p17585144585310"><a name="p17585144585310"></a><a name="p17585144585310"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p10783112824619"><a name="p10783112824619"></a><a name="p10783112824619"></a>Mandatory for downloading an object</p>
<p id="p19201728661"><a name="p19201728661"></a><a name="p19201728661"></a>Optional for downloading objects in a batch</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p0588134575318"><a name="p0588134575318"></a><a name="p0588134575318"></a>Indicates the name of the object to be downloaded, or the name prefix of the objects to be downloaded in batches.</p>
<p id="p10591154565315"><a name="p10591154565315"></a><a name="p10591154565315"></a>This parameter cannot be left blank when downloading an object. The saving and naming rules are as follows:</p>
<a name="ul105928456532"></a><a name="ul105928456532"></a><ul id="ul105928456532"><li>If this parameter specifies a file or folder path that does not exist, the tool checks whether the value ends with a slash (/) or backslash (\). If yes, a folder is created based on the path, and the object is downloaded to this newly created directory.</li><li>If this parameter specifies a file or folder path that does not exist and the value does not end with a slash (/) or backslash (\), the object is downloaded to your local PC with the value of <strong id="b39171932103717"><a name="b39171932103717"></a><a name="b39171932103717"></a>key</strong> as the file name.</li><li>If this parameter specifies an existing file, the object is downloaded to your local PC overwriting the existing file, with the value of <strong id="b1581145618414"><a name="b1581145618414"></a><a name="b1581145618414"></a>key</strong> as the file name.</li><li>If this parameter specifies an existing folder, the object is downloaded to the directory specified by <strong id="b18142520175817"><a name="b18142520175817"></a><a name="b18142520175817"></a>file_or_folder_url</strong> with the object name as the file name.</li></ul>
<p id="p115612190125"><a name="p115612190125"></a><a name="p115612190125"></a>The saving rules when downloading objects in batches are as follows:</p>
<a name="ul15158191991211"></a><a name="ul15158191991211"></a><ul id="ul15158191991211"><li>If this parameter is left blank, all objects in the bucket are downloaded to the directory specified by <strong id="b19288175202813"><a name="b19288175202813"></a><a name="b19288175202813"></a>folder_url</strong>.</li><li>If this parameter is configured, objects whose name prefix is the configured value in the bucket are downloaded to the directory specified by <strong id="b202481059123017"><a name="b202481059123017"></a><a name="b202481059123017"></a>folder_url</strong>.</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul54182557214"></a><a name="ul54182557214"></a><ul id="ul54182557214"><li>If this parameter is configured but the <strong id="b25692434515"><a name="b25692434515"></a><a name="b25692434515"></a>flat</strong> parameter is not configured when downloading objects in batches, the name of the downloaded file contains the name prefix of the parent object. If <strong id="b80195614562"><a name="b80195614562"></a><a name="b80195614562"></a>flat</strong> is configured, then the name of the downloaded file does not contain the name prefix of the parent object.</li><li>For details about how to use this parameter, see <a href="download-examples.md">Download Examples</a>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1314444315116"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p123051451513"><a name="p123051451513"></a><a name="p123051451513"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p113071451811"><a name="p113071451811"></a><a name="p113071451811"></a>Optional for downloading an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p73101351816"><a name="p73101351816"></a><a name="p73101351816"></a>Generates an operation result list when downloading an object.</p>
</td>
</tr>
<tr id="row20942144811418"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p1633845116148"><a name="p1633845116148"></a><a name="p1633845116148"></a>flat</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p43401151111418"><a name="p43401151111418"></a><a name="p43401151111418"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p2343205119145"><a name="p2343205119145"></a><a name="p2343205119145"></a>The name prefix of the parent object is excluded when downloading objects in batches.</p>
</td>
</tr>
<tr id="row1163113812718"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p276284012715"><a name="p276284012715"></a><a name="p276284012715"></a>tempFileDir</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p122055912103"><a name="p122055912103"></a><a name="p122055912103"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1476264014710"><a name="p1476264014710"></a><a name="p1476264014710"></a>Indicates the directory for storing temporary files during multipart download. The default value is the value of <strong id="b19183455528"><a name="b19183455528"></a><a name="b19183455528"></a>defaultTempFileDir</strong> in the configuration file.</p>
<div class="note" id="note18655347142618"><a name="note18655347142618"></a><a name="note18655347142618"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul15994748142619"></a><a name="ul15994748142619"></a><ul id="ul15994748142619"><li>If this parameter is left blank and the <strong id="b1049928913"><a name="b1049928913"></a><a name="b1049928913"></a>defaultTempFileDir</strong> parameter in the configuration file is also left blank, temporary files generated during multipart download are saved in the directory where to-be-downloaded files are located and end with the suffix of <strong id="b1054316270311"><a name="b1054316270311"></a><a name="b1054316270311"></a>.obs.temp</strong>.</li><li>Temporary files generated during multipart download are stored in this directory. Therefore, ensure that the user who executes obsutil has the write permission on the path.</li><li>The available space of the partition where the path is located must be greater than the size of the objects to be downloaded.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1525855205616"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row35998457538"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p206001459530"><a name="p206001459530"></a><a name="p206001459530"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p10601174595317"><a name="p10601174595317"></a><a name="p10601174595317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1260204518534"><a name="p1260204518534"></a><a name="p1260204518534"></a>Indicates incremental download. If this parameter is set, each object can be downloaded only when it does not exist in the local path, its size is different from the namesake one in the local path, or it has the latest modification time.</p>
</td>
</tr>
<tr id="row9603194595311"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p1605745105311"><a name="p1605745105311"></a><a name="p1605745105311"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p12605114512534"><a name="p12605114512534"></a><a name="p12605114512534"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1660724595316"><a name="p1660724595316"></a><a name="p1660724595316"></a>Checks whether the sizes of the local files are the same as those of the objects in the bucket after the download is complete.</p>
</td>
</tr>
<tr id="row360812456535"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p136091945105315"><a name="p136091945105315"></a><a name="p136091945105315"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p461011452539"><a name="p461011452539"></a><a name="p461011452539"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p19611745185310"><a name="p19611745185310"></a><a name="p19611745185310"></a>Checks whether MD5 values of the local files are the same as those of the objects in the bucket after the download is complete.</p>
<div class="note" id="note2303131110278"><a name="note2303131110278"></a><a name="note2303131110278"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p12303311112719"><a name="p12303311112719"></a><a name="p12303311112719"></a>Objects in the bucket must contain metadata <strong id="b1409155717507"><a name="b1409155717507"></a><a name="b1409155717507"></a>x-obs-md5chksum</strong>. Otherwise, MD5 verification will be skipped.</p>
</div></div>
</td>
</tr>
<tr id="row36131445115310"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p14613194514535"><a name="p14613194514535"></a><a name="p14613194514535"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p15615164515317"><a name="p15615164515317"></a><a name="p15615164515317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1161744513534"><a name="p1161744513534"></a><a name="p1161744513534"></a>Indicates the maximum number of concurrent multipart download tasks when downloading an object. The default value is the value of <strong id="b18251123671818"><a name="b18251123671818"></a><a name="b18251123671818"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row14617144520537"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p17619545185314"><a name="p17619545185314"></a><a name="p17619545185314"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p9620645185311"><a name="p9620645185311"></a><a name="p9620645185311"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p262174575315"><a name="p262174575315"></a><a name="p262174575315"></a>Indicates the threshold for enabling multipart download, in bytes. The default value is the value of <strong id="b86361877420"><a name="b86361877420"></a><a name="b86361877420"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note94444118613"><a name="note94444118613"></a><a name="note94444118613"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be downloaded is smaller than the threshold, download the object directly. If not, a multipart download is required.</li><li>If you download an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b13882132920443"><a name="b13882132920443"></a><a name="b13882132920443"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row136222045105310"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p122913325265"><a name="p122913325265"></a><a name="p122913325265"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p1662574595317"><a name="p1662574595317"></a><a name="p1662574595317"></a>Optional for downloading an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p662724555310"><a name="p662724555310"></a><a name="p662724555310"></a>Source object version ID that can be specified when downloading an object</p>
</td>
</tr>
<tr id="row1363634595315"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p3637164545318"><a name="p3637164545318"></a><a name="p3637164545318"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p1264164517532"><a name="p1264164517532"></a><a name="p1264164517532"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p10641184516531"><a name="p10641184516531"></a><a name="p10641184516531"></a>Indicates the size of each part in a multipart download task, in bytes. The default value is the value of <strong id="b1530220454436"><a name="b1530220454436"></a><a name="b1530220454436"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note1175045524310"><a name="note1175045524310"></a><a name="note1175045524310"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul230316596421"></a><a name="ul230316596421"></a><ul id="ul230316596421"><li>This value can contain a capacity unit. For example, <strong id="b1776594514438"><a name="b1776594514438"></a><a name="b1776594514438"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b5916847154318"><a name="b5916847154318"></a><a name="b5916847154318"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row196461145155314"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p1764874511530"><a name="p1764874511530"></a><a name="p1764874511530"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p206482455537"><a name="p206482455537"></a><a name="p206482455537"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1965018458536"><a name="p1965018458536"></a><a name="p1965018458536"></a>Indicates the folder where the part records reside. The default value is <strong id="b1239112918205"><a name="b1239112918205"></a><a name="b1239112918205"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note519835512711"><a name="note519835512711"></a><a name="note519835512711"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart download and saved to the <strong id="b1575171113474"><a name="b1575171113474"></a><a name="b1575171113474"></a>down</strong> subfolder. After the download succeeds, its part record is deleted automatically. If the download fails or is suspended, the system attempts to resume the task according to its part record when you perform the download the next time.</p>
</div></div>
</td>
</tr>
<tr id="row1165094575317"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p96521245195311"><a name="p96521245195311"></a><a name="p96521245195311"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p7654154512534"><a name="p7654154512534"></a><a name="p7654154512534"></a>Mandatory for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p206552045145317"><a name="p206552045145317"></a><a name="p206552045145317"></a>Copies objects in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row5655184520536"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p56577450535"><a name="p56577450535"></a><a name="p56577450535"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p13658645125314"><a name="p13658645125314"></a><a name="p13658645125314"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p566014515533"><a name="p566014515533"></a><a name="p566014515533"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1566014510533"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p3662204585313"><a name="p3662204585313"></a><a name="p3662204585313"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p19305175612615"><a name="p19305175612615"></a><a name="p19305175612615"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p17665114555311"><a name="p17665114555311"></a><a name="p17665114555311"></a>Indicates the maximum number of concurrent tasks for downloading objects in a batch. The default value is the value of <strong id="b27577263238"><a name="b27577263238"></a><a name="b27577263238"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row109762261097"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p156173113915"><a name="p156173113915"></a><a name="p156173113915"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p1930885652611"><a name="p1930885652611"></a><a name="p1930885652611"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b14546165310472"><a name="b14546165310472"></a><a name="b14546165310472"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b68461455154720"><a name="b68461455154720"></a><a name="b68461455154720"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b0847175564713"><a name="b0847175564713"></a><a name="b0847175564713"></a>abc</strong> and ends with <strong id="b11848125564714"><a name="b11848125564714"></a><a name="b11848125564714"></a>.txt</strong>.</li><li>You can use <strong id="b1894173812108"><a name="b1894173812108"></a><a name="b1894173812108"></a>\*</strong> to represent <strong id="b689573881013"><a name="b689573881013"></a><a name="b689573881013"></a>*</strong> and <strong id="b2089518388106"><a name="b2089518388106"></a><a name="b2089518388106"></a>\?</strong> to represent <strong id="b1789683815103"><a name="b1789683815103"></a><a name="b1789683815103"></a>?</strong>.</li><li>If the name of the object to be downloaded matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b8572151512274"><a name="b8572151512274"></a><a name="b8572151512274"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b9573115102713"><a name="b9573115102713"></a><a name="b9573115102713"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row20666154513530"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p186661845145315"><a name="p186661845145315"></a><a name="p186661845145315"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p966784516533"><a name="p966784516533"></a><a name="p966784516533"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b48617711102"><a name="b48617711102"></a><a name="b48617711102"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b96364102108"><a name="b96364102108"></a><a name="b96364102108"></a>\*</strong> to represent <strong id="b136376101104"><a name="b136376101104"></a><a name="b136376101104"></a>*</strong> and <strong id="b2638710121019"><a name="b2638710121019"></a><a name="b2638710121019"></a>\?</strong> to represent <strong id="b1263991081017"><a name="b1263991081017"></a><a name="b1263991081017"></a>?</strong>.</li><li>Only after identifying that the name of the file to be downloaded does not match the value of <strong id="b14699101251018"><a name="b14699101251018"></a><a name="b14699101251018"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is downloaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul1329413124584"></a><a name="ul1329413124584"></a><ul id="ul1329413124584"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b14412319182712"><a name="b14412319182712"></a><a name="b14412319182712"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b14412121942719"><a name="b14412121942719"></a><a name="b14412121942719"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row949610111431"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when downloading objects. Only objects whose last modification time is within the configured time range are downloaded.</p>
<p id="p338461155315"><a name="p338461155315"></a><a name="p338461155315"></a>This pattern has a lower priority than the object matching patterns (<strong id="b239512331936"><a name="b239512331936"></a><a name="b239512331936"></a>exclude</strong>/<strong id="b33978336316"><a name="b33978336316"></a><a name="b33978336316"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i7149142272719"><a name="i7149142272719"></a><a name="i7149142272719"></a>time1</em><strong id="b17149422192717"><a name="b17149422192717"></a><a name="b17149422192717"></a>-</strong><em id="i8149162242711"><a name="i8149162242711"></a><a name="i8149162242711"></a>time2</em>, where <em id="i61506224276"><a name="i61506224276"></a><a name="i61506224276"></a>time1</em> must be earlier than or the same as <em id="i201501622202716"><a name="i201501622202716"></a><a name="i201501622202716"></a>time2</em>. The time format is <em id="i4151722102717"><a name="i4151722102717"></a><a name="i4151722102717"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b1156152520273"><a name="b1156152520273"></a><a name="b1156152520273"></a>*-</strong><em id="i856132562720"><a name="i856132562720"></a><a name="i856132562720"></a>time2</em>, all files whose last modification time is earlier than <em id="i857132511273"><a name="i857132511273"></a><a name="i857132511273"></a>time2</em> are matched. If it is set to <em id="i205812582712"><a name="i205812582712"></a><a name="i205812582712"></a>time1</em><strong id="b559225112719"><a name="b559225112719"></a><a name="b559225112719"></a>-*</strong>, all files whose last modification time is later than <em id="i15608259273"><a name="i15608259273"></a><a name="i15608259273"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row1116705416719"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b4613559637"><a name="b4613559637"></a><a name="b4613559637"></a>include</strong> or <strong id="b561319593313"><a name="b561319593313"></a><a name="b561319593313"></a>exclude</strong>) and the time matching pattern (<strong id="b86141759132"><a name="b86141759132"></a><a name="b86141759132"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row109093549288"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p143226572287"><a name="p143226572287"></a><a name="p143226572287"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p232405732811"><a name="p232405732811"></a><a name="p232405732811"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p1632615718282"><a name="p1632615718282"></a><a name="p1632615718282"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b716416162111"><a name="b716416162111"></a><a name="b716416162111"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note1416817409247"><a name="note1416817409247"></a><a name="note1416817409247"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b8752111795813"><a name="b8752111795813"></a><a name="b8752111795813"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i775314172581"><a name="i775314172581"></a><a name="i775314172581"></a>time</em><strong id="b5754917135818"><a name="b5754917135818"></a><a name="b5754917135818"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b292825133616"><a name="b292825133616"></a><a name="b292825133616"></a>recordMaxLogSize</strong> and <strong id="b179352516361"><a name="b179352516361"></a><a name="b179352516361"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1920762633614"><td class="cellrowborder" valign="top" width="17.17171717171717%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="26.26262626262626%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.56565656565656%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section1059235632510"></a>

-   Take the Windows OS as an example. Run the  **obsutil cp obs://bucket-test/key** **d:\\temp\\test.txt**  command to download a single object.

```
obsutil cp obs://bucket-test/key d:\temp\test.txt

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx

[==========================================] 100.00% 4.86 KB/s 0s
Download successfully, 19B, obs://bucket-test/key --> d:\temp\test.txt
```

-   Take the Windows OS as an example. Run the  **obsutil cp obs://bucket-test/temp d:\\ -f -r**  command to download objects in batches.

```
obsutil cp obs://bucket-test/temp d:\ -f -r

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx

[======================================================] 100.00% 155.59 KB/s 0s
Succeed count is:   6         Failed count is:    0
Metrics [max cost:153 ms, min cost:129 ms, average cost:92.00 ms, average tps:17.86]
Task id is: 3066a4b0-4d21-4929-bb84-4829c32cbd0f
```

-   For more examples, see  [Download Examples](download-examples.md).

