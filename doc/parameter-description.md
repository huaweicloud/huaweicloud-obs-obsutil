# Parameter Description<a name="EN-US_TOPIC_0147305055"></a>

You can use the  **.obsutilconfig**  file to configure the parameters of obsutil. The following table lists detailed information about the parameters.

**Table  1**  obsutil parameters

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="15%" id="mcps1.2.5.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="18%" id="mcps1.2.5.1.2"><p id="p11492115213118"><a name="p11492115213118"></a><a name="p11492115213118"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="42%" id="mcps1.2.5.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
<th class="cellrowborder" valign="top" width="25%" id="mcps1.2.5.1.4"><p id="p174671008433"><a name="p174671008433"></a><a name="p174671008433"></a>Recommended Value</p>
</th>
</tr>
</thead>
<tbody><tr id="row1371911524145"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p89611756121419"><a name="p89611756121419"></a><a name="p89611756121419"></a>endpoint</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1796311568142"><a name="p1796311568142"></a><a name="p1796311568142"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p396811562144"><a name="p396811562144"></a><a name="p396811562144"></a>Endpoint for accessing OBS, which can contain the protocol type, domain name, and port number (optional). Example: <strong id="b1944782065311"><a name="b1944782065311"></a><a name="b1944782065311"></a>https://your-endpoint:80</strong></p>
<div class="note" id="note19977155618144"><a name="note19977155618144"></a><a name="note19977155618144"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul14981256161413"></a><a name="ul14981256161413"></a><ul id="ul14981256161413"><li>You can click <a href="https://docs.otc.t-systems.com/en-us/endpoint/index.html" target="_blank" rel="noopener noreferrer">here</a> to view the endpoints and regions enabled for OBS.</li><li>If the configured endpoint does not contain any protocol, the HTTPS protocol is used by default.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p2986956101417"><a name="p2986956101417"></a><a name="p2986956101417"></a>N/A</p>
</td>
</tr>
<tr id="row108328217449"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p16832192117441"><a name="p16832192117441"></a><a name="p16832192117441"></a>ak</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p8494852131113"><a name="p8494852131113"></a><a name="p8494852131113"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p6668151125217"><a name="p6668151125217"></a><a name="p6668151125217"></a>Access key ID</p>
<div class="note" id="note166811942674"><a name="note166811942674"></a><a name="note166811942674"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul17232014104512"></a><a name="ul17232014104512"></a><ul id="ul17232014104512"><li>After you run obsutil for the first time, the tool encrypts the AK to ensure the key security.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p16187131231020"><a name="p16187131231020"></a><a name="p16187131231020"></a>N/A</p>
</td>
</tr>
<tr id="row48331421104417"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p18833142116449"><a name="p18833142116449"></a><a name="p18833142116449"></a>sk</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1549410520111"><a name="p1549410520111"></a><a name="p1549410520111"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p965412289498"><a name="p965412289498"></a><a name="p965412289498"></a>Secret access key</p>
<div class="note" id="note1533418277817"><a name="note1533418277817"></a><a name="note1533418277817"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul20420183804419"></a><a name="ul20420183804419"></a><ul id="ul20420183804419"><li>After you run obsutil for the first time, the tool encrypts the SK to ensure the key security.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p321061211103"><a name="p321061211103"></a><a name="p321061211103"></a>N/A</p>
</td>
</tr>
<tr id="row0838143916237"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p183818392232"><a name="p183818392232"></a><a name="p183818392232"></a>endpointCrr</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p19581786246"><a name="p19581786246"></a><a name="p19581786246"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p352129122619"><a name="p352129122619"></a><a name="p352129122619"></a>Endpoint for accessing OBS in the region where the source bucket resides when the client-side cross-region replication function is enabled, which can contain the protocol type, domain name, and port number. Example: <strong id="b1273285569"><a name="b1273285569"></a><a name="b1273285569"></a>http://your-endpoint:80</strong></p>
<div class="note" id="note85252932620"><a name="note85252932620"></a><a name="note85252932620"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul1751945114417"></a><a name="ul1751945114417"></a>
<a name="ul175252914269"></a><a name="ul175252914269"></a><ul id="ul175252914269"><li>You can click <a href="https://docs.otc.t-systems.com/en-us/endpoint/index.html" target="_blank" rel="noopener noreferrer">here</a> to view the endpoints and regions enabled for OBS.</li><li>If the configured endpoint does not contain any protocol, the HTTPS protocol is used by default.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p5870014122419"><a name="p5870014122419"></a><a name="p5870014122419"></a>N/A</p>
</td>
</tr>
<tr id="row137327363233"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p57331336162316"><a name="p57331336162316"></a><a name="p57331336162316"></a>akCrr</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1659611842417"><a name="p1659611842417"></a><a name="p1659611842417"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p1253529102619"><a name="p1253529102619"></a><a name="p1253529102619"></a>AK for the source bucket when the client-side cross-region replication function is enabled</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p158811214162416"><a name="p158811214162416"></a><a name="p158811214162416"></a>N/A</p>
</td>
</tr>
<tr id="row134663317238"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p10347153312231"><a name="p10347153312231"></a><a name="p10347153312231"></a>skCrr</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p126075802412"><a name="p126075802412"></a><a name="p126075802412"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p11531629152612"><a name="p11531629152612"></a><a name="p11531629152612"></a>SK for the source bucket when the client-side cross-region replication function is enabled</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p138911014182415"><a name="p138911014182415"></a><a name="p138911014182415"></a>N/A</p>
</td>
</tr>
<tr id="row116183195211"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1316114335213"><a name="p1316114335213"></a><a name="p1316114335213"></a>token</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p54941952201116"><a name="p54941952201116"></a><a name="p54941952201116"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p1666201165210"><a name="p1666201165210"></a><a name="p1666201165210"></a>Security token. If this parameter is left blank, the security token is not set.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p8214111210109"><a name="p8214111210109"></a><a name="p8214111210109"></a>N/A</p>
</td>
</tr>
<tr id="row1229802017353"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p2298142013513"><a name="p2298142013513"></a><a name="p2298142013513"></a>connectTimeout</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p5494135214110"><a name="p5494135214110"></a><a name="p5494135214110"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p02993209355"><a name="p02993209355"></a><a name="p02993209355"></a>Timeout interval for establishing an HTTP/HTTPS connection, in seconds. The default value is 30.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p18299132019354"><a name="p18299132019354"></a><a name="p18299132019354"></a>The recommended value ranges from 5 to 120.</p>
</td>
</tr>
<tr id="row116081618193518"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p116099188354"><a name="p116099188354"></a><a name="p116099188354"></a>socketTimeout</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p204948520113"><a name="p204948520113"></a><a name="p204948520113"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p16092018143515"><a name="p16092018143515"></a><a name="p16092018143515"></a>Timeout interval for reading and writing data, in seconds. The default value is 310.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p17609131813519"><a name="p17609131813519"></a><a name="p17609131813519"></a>The recommended value ranges from 5 to 600.</p>
</td>
</tr>
<tr id="row145700495434"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p14572549174310"><a name="p14572549174310"></a><a name="p14572549174310"></a>maxRetryCount</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p94941652171112"><a name="p94941652171112"></a><a name="p94941652171112"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p1857224911439"><a name="p1857224911439"></a><a name="p1857224911439"></a>Maximum number of retry attempts. The default value is 3.</p>
<div class="note" id="note1518156154417"><a name="note1518156154417"></a><a name="note1518156154417"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p626795614461"><a name="p626795614461"></a><a name="p626795614461"></a>When an OBS request completes but HTTP status code 408 or 5XX is returned, or when a timeout error occurs to an OBS request, the request is retried.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p2572194912431"><a name="p2572194912431"></a><a name="p2572194912431"></a>The recommended value ranges from 0 to 5.</p>
</td>
</tr>
<tr id="row1251775818314"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1519185819313"><a name="p1519185819313"></a><a name="p1519185819313"></a>maxConnections</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p14519175863112"><a name="p14519175863112"></a><a name="p14519175863112"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p983302124412"><a name="p983302124412"></a><a name="p983302124412"></a>Maximum number of HTTP connections that can be accessed. The default value is 1000.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1751955818315"><a name="p1751955818315"></a><a name="p1751955818315"></a>N/A</p>
</td>
</tr>
<tr id="row6502192717387"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1050292773816"><a name="p1050292773816"></a><a name="p1050292773816"></a>defaultBigfileThreshold</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p114941252171115"><a name="p114941252171115"></a><a name="p114941252171115"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p9662911185214"><a name="p9662911185214"></a><a name="p9662911185214"></a>Indicates the threshold for triggering multipart tasks, in bytes. If the size of a file to be uploaded, downloaded, or copied is greater than the threshold, the file is uploaded, downloaded, or copied in multiple parts. The default value is 50 MB.</p>
<div class="note" id="note118991088514"><a name="note118991088514"></a><a name="note118991088514"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p29001283516"><a name="p29001283516"></a><a name="p29001283516"></a>This value can contain a capacity unit. For example, <strong id="b1224933118515"><a name="b1224933118515"></a><a name="b1224933118515"></a>1 MB</strong> indicates 1048576 bytes.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1166191125217"><a name="p1166191125217"></a><a name="p1166191125217"></a>It is recommended that the value be greater than 5 MB.</p>
</td>
</tr>
<tr id="row1577814816513"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p077914845118"><a name="p077914845118"></a><a name="p077914845118"></a>defaultPartSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p11494052181113"><a name="p11494052181113"></a><a name="p11494052181113"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p161953211218"><a name="p161953211218"></a><a name="p161953211218"></a>Size of each part, in bytes. The default value is <strong id="b79039391754"><a name="b79039391754"></a><a name="b79039391754"></a>auto</strong>.</p>
<div class="note" id="note14967102951311"><a name="note14967102951311"></a><a name="note14967102951311"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul14427184914139"></a><a name="ul14427184914139"></a><ul id="ul14427184914139"><li>For multipart upload and copy, the value ranges from 100 KB to 5 GB.</li><li>For multipart download, the value is unrestricted.</li><li>This value can contain a capacity unit. For example, <strong id="b1626111469511"><a name="b1626111469511"></a><a name="b1626111469511"></a>1 MB</strong> indicates 1048576 bytes.</li><li>If this parameter is set to <strong id="b6905144713512"><a name="b6905144713512"></a><a name="b6905144713512"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p465918118528"><a name="p465918118528"></a><a name="p465918118528"></a>[9MB, 100MB]</p>
</td>
</tr>
<tr id="row168335218440"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p4866163316380"><a name="p4866163316380"></a><a name="p4866163316380"></a>defaultParallels</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p17494135210112"><a name="p17494135210112"></a><a name="p17494135210112"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p986818333384"><a name="p986818333384"></a><a name="p986818333384"></a>Maximum number of concurrent tasks in the multipart mode. The default value is 5.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1086920338382"><a name="p1086920338382"></a><a name="p1086920338382"></a>Set this parameter according to <a href="fine-tuning-obsutil-performance.md">Fine-Tuning obsutil Performance</a>.</p>
</td>
</tr>
<tr id="row1272522783817"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1126123113812"><a name="p1126123113812"></a><a name="p1126123113812"></a>defaultJobs</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p849417529115"><a name="p849417529115"></a><a name="p849417529115"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p112661531143816"><a name="p112661531143816"></a><a name="p112661531143816"></a>Maximum number of concurrent tasks in batches. The default value is 5.</p>
<div class="note" id="note174841782158"><a name="note174841782158"></a><a name="note174841782158"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p44851810155"><a name="p44851810155"></a><a name="p44851810155"></a>Batch tasks include uploading, downloading, and copying folders, as well as restoring and deleting objects in batches.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1826743173815"><a name="p1826743173815"></a><a name="p1826743173815"></a>[1, 50]</p>
</td>
</tr>
<tr id="row15833621134411"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p68331421164412"><a name="p68331421164412"></a><a name="p68331421164412"></a>defaultJobsCacheCount</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p144949527111"><a name="p144949527111"></a><a name="p144949527111"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p965571185218"><a name="p965571185218"></a><a name="p965571185218"></a>Cache size of a batch task queue, indicating the maximum number of tasks that can be cached. The default value is 1000000.</p>
<div class="note" id="note01031142172"><a name="note01031142172"></a><a name="note01031142172"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p20685192420498"><a name="p20685192420498"></a><a name="p20685192420498"></a>More cached tasks consume more memory resources. Therefore, you are advised to adjust the value of this parameter based on site requirements.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p138341128183911"><a name="p138341128183911"></a><a name="p138341128183911"></a>Default value</p>
</td>
</tr>
<tr id="row18901163111115"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p19012313115"><a name="p19012313115"></a><a name="p19012313115"></a>rateLimitThreshold</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p6494252111116"><a name="p6494252111116"></a><a name="p6494252111116"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p11653151115219"><a name="p11653151115219"></a><a name="p11653151115219"></a>Indicates the traffic control threshold of an upload or download request, in bytes. The default value is 0, indicating that traffic is not limited. The minimum value is 10 KB.</p>
<div class="note" id="note9215194216506"><a name="note9215194216506"></a><a name="note9215194216506"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p8216124265017"><a name="p8216124265017"></a><a name="p8216124265017"></a>This value can contain a capacity unit. For example, <strong id="b243434110610"><a name="b243434110610"></a><a name="b243434110610"></a>1 MB</strong> indicates 1048576 bytes.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1420814163918"><a name="p1420814163918"></a><a name="p1420814163918"></a>It is recommended that the value be greater than 100 KB.</p>
</td>
</tr>
<tr id="row0600952111317"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p4600552171310"><a name="p4600552171310"></a><a name="p4600552171310"></a>sdkLogBackups</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p14494135251111"><a name="p14494135251111"></a><a name="p14494135251111"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p6757235"><a name="p6757235"></a><a name="p6757235"></a>Maximum number of SDK log files that can be retained. The default value is 10.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1629625292219"><a name="p1629625292219"></a><a name="p1629625292219"></a>N/A</p>
</td>
</tr>
<tr id="row1376292910153"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1176217292152"><a name="p1176217292152"></a><a name="p1176217292152"></a>sdkLogLevel</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p7494175211119"><a name="p7494175211119"></a><a name="p7494175211119"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p28644761"><a name="p28644761"></a><a name="p28644761"></a>SDK log level. Possible values are:</p>
<a name="ul1654371432110"></a><a name="ul1654371432110"></a><ul id="ul1654371432110"><li>DEBUG</li><li>INFO</li><li>WARN</li><li>ERROR</li></ul>
<p id="p849741142512"><a name="p849741142512"></a><a name="p849741142512"></a>The default value is <strong id="b1634326835"><a name="b1634326835"></a><a name="b1634326835"></a>WARN</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1330015523223"><a name="p1330015523223"></a><a name="p1330015523223"></a>N/A</p>
</td>
</tr>
<tr id="row106540991712"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1065417941713"><a name="p1065417941713"></a><a name="p1065417941713"></a>sdkLogPath</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p6494452121111"><a name="p6494452121111"></a><a name="p6494452121111"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p4649191145220"><a name="p4649191145220"></a><a name="p4649191145220"></a>Indicates the absolute path of SDK logs. The value must be a file path. The default value is the path of the <strong id="b67597401482"><a name="b67597401482"></a><a name="b67597401482"></a>obssdk.log</strong> file in the subfolder <strong id="b20755511492"><a name="b20755511492"></a><a name="b20755511492"></a>obsutil_log</strong> of the user's home directory (<strong id="b669212479115"><a name="b669212479115"></a><a name="b669212479115"></a>HOME</strong> in Linux or macOS and <strong id="b3605145141210"><a name="b3605145141210"></a><a name="b3605145141210"></a>C:\Users\</strong><em id="i395564810127"><a name="i395564810127"></a><a name="i395564810127"></a>&lt;Username&gt;</em> in Windows).</p>
<div class="note" id="note19394054141910"><a name="note19394054141910"></a><a name="note19394054141910"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul897844719126"></a><a name="ul897844719126"></a><ul id="ul897844719126"><li>If this parameter is left blank, no SDK log is generated.</li><li>The path must be a file path and cannot be a folder path.</li><li>After the SDK log function is enabled, all logs of requests to OBS are saved in the SDK log file for problem analysis and location.</li><li>Ensure that the user who runs the command has the read and write permissions on the path.</li></ul>
</div></div>
<div class="notice" id="note022594219386"><a name="note022594219386"></a><a name="note022594219386"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p13226164214389"><a name="p13226164214389"></a><a name="p13226164214389"></a>If multiple obsutil processes are running at the same time, log files may fail to be written concurrently or may be lost. In this case, add parameter <strong id="b184878130229"><a name="b184878130229"></a><a name="b184878130229"></a>-config</strong> when running commands to configure an independent configuration file for each process.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p143028525228"><a name="p143028525228"></a><a name="p143028525228"></a>N/A</p>
</td>
</tr>
<tr id="row133933491190"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p16393124951913"><a name="p16393124951913"></a><a name="p16393124951913"></a>sdkMaxLogSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p0494185261114"><a name="p0494185261114"></a><a name="p0494185261114"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p5646811125218"><a name="p5646811125218"></a><a name="p5646811125218"></a>Size of an SDK log file, in bytes. The default value is 30 MB.</p>
<div class="note" id="note34732361504"><a name="note34732361504"></a><a name="note34732361504"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p6474133616508"><a name="p6474133616508"></a><a name="p6474133616508"></a>This value can contain a capacity unit. For example, <strong id="b11228204815618"><a name="b11228204815618"></a><a name="b11228204815618"></a>1 MB</strong> indicates 1048576 bytes.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p8305185282215"><a name="p8305185282215"></a><a name="p8305185282215"></a>The recommended value ranges from 10 MB to 100 MB.</p>
</td>
</tr>
<tr id="row19437155319193"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p17437165381918"><a name="p17437165381918"></a><a name="p17437165381918"></a>utilLogBackups</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p34941252201112"><a name="p34941252201112"></a><a name="p34941252201112"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p390915288213"><a name="p390915288213"></a><a name="p390915288213"></a>Maximum number of obsutil log files that can be retained. The default value is 10.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p53081852142211"><a name="p53081852142211"></a><a name="p53081852142211"></a>N/A</p>
</td>
</tr>
<tr id="row8668194319411"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p2422925171916"><a name="p2422925171916"></a><a name="p2422925171916"></a>utilLogLevel</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p0494195221116"><a name="p0494195221116"></a><a name="p0494195221116"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p917204152114"><a name="p917204152114"></a><a name="p917204152114"></a>obsutil log level. Possible values are:</p>
<a name="ul317204113212"></a><a name="ul317204113212"></a><ul id="ul317204113212"><li>DEBUG</li><li>INFO</li><li>WARN</li><li>ERROR</li></ul>
<p id="p16952155314259"><a name="p16952155314259"></a><a name="p16952155314259"></a>The default value is <strong id="b1366843013267"><a name="b1366843013267"></a><a name="b1366843013267"></a>INFO</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p14311185222214"><a name="p14311185222214"></a><a name="p14311185222214"></a>N/A</p>
</td>
</tr>
<tr id="row756217410416"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p10703134581918"><a name="p10703134581918"></a><a name="p10703134581918"></a>utilLogPath</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p649419526116"><a name="p649419526116"></a><a name="p649419526116"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p185241522185012"><a name="p185241522185012"></a><a name="p185241522185012"></a>Indicates the absolute path of obsutil logs. The value must be a file path. The default value is the path of the <strong id="b164821925161310"><a name="b164821925161310"></a><a name="b164821925161310"></a>obsutil.log</strong> file in the subfolder <strong id="b64831625141316"><a name="b64831625141316"></a><a name="b64831625141316"></a>.obsutil_log</strong> of the user's home directory (<strong id="b948413257137"><a name="b948413257137"></a><a name="b948413257137"></a>HOME</strong> in Linux or macOS and <strong id="b9485825161311"><a name="b9485825161311"></a><a name="b9485825161311"></a>C:\Users\</strong><em id="i348552514132"><a name="i348552514132"></a><a name="i348552514132"></a>&lt;Username&gt;</em> in Windows).</p>
<div class="note" id="note853813419401"><a name="note853813419401"></a><a name="note853813419401"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19891125164011"></a><a name="ul19891125164011"></a><ul id="ul19891125164011"><li>If this parameter is left blank, no obsutil log is generated.</li><li>The path must be a file path and cannot be a folder path.</li><li>After the obsutil log function is enabled, all logs generated during commands executing are saved in the obsutil log file for problem analysis and location.</li><li>Ensure that the user who runs the command has the read and write permissions on the path.</li></ul>
</div></div>
<div class="notice" id="note2669149184013"><a name="note2669149184013"></a><a name="note2669149184013"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p146691794409"><a name="p146691794409"></a><a name="p146691794409"></a>If multiple obsutil processes are running at the same time, log files may fail to be written concurrently. In this case, add parameter <strong id="b88101327217"><a name="b88101327217"></a><a name="b88101327217"></a>-config</strong> when running commands to configure an independent configuration file for each process.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p03143522222"><a name="p03143522222"></a><a name="p03143522222"></a>N/A</p>
</td>
</tr>
<tr id="row351654320425"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p9150451113118"><a name="p9150451113118"></a><a name="p9150451113118"></a>utilMaxLogSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p74946525116"><a name="p74946525116"></a><a name="p74946525116"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p49291135162216"><a name="p49291135162216"></a><a name="p49291135162216"></a>Size of an obsutil log file, in bytes. The default value is 30 MB.</p>
<div class="note" id="note121371115165012"><a name="note121371115165012"></a><a name="note121371115165012"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1813819153502"><a name="p1813819153502"></a><a name="p1813819153502"></a>This value can contain a capacity unit. For example, <strong id="b1472613541865"><a name="b1472613541865"></a><a name="b1472613541865"></a>1 MB</strong> indicates 1048576 bytes.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1545515615451"><a name="p1545515615451"></a><a name="p1545515615451"></a>The recommended value ranges from 10 MB to 100 MB.</p>
</td>
</tr>
<tr id="row767611194491"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p9107202294911"><a name="p9107202294911"></a><a name="p9107202294911"></a>writeBufferIoSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1610922294917"><a name="p1610922294917"></a><a name="p1610922294917"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p611012284912"><a name="p611012284912"></a><a name="p611012284912"></a>Size of the cache for downloading data, in bytes. The default value is 65536.</p>
<div class="note" id="note1511342214917"><a name="note1511342214917"></a><a name="note1511342214917"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul661854710521"></a><a name="ul661854710521"></a><ul id="ul661854710521"><li>Set this parameter based on site requirements. If the size of the file to be downloaded is large, set this parameter to a large value.</li><li>This value can contain a capacity unit. For example, <strong id="b5930235714"><a name="b5930235714"></a><a name="b5930235714"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p13121922204918"><a name="p13121922204918"></a><a name="p13121922204918"></a>N/A</p>
</td>
</tr>
<tr id="row1297614241348"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p12977192413417"><a name="p12977192413417"></a><a name="p12977192413417"></a>readBufferIoSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p858121153519"><a name="p858121153519"></a><a name="p858121153519"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p1104161214358"><a name="p1104161214358"></a><a name="p1104161214358"></a>Size of the cache for uploading data, in bytes. The default value is 8192.</p>
<div class="note" id="note4104192363619"><a name="note4104192363619"></a><a name="note4104192363619"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul113272040125215"></a><a name="ul113272040125215"></a><ul id="ul113272040125215"><li>Set this parameter based on site requirements. If a large number of small files are uploaded, set this parameter to a small value. If large files are uploaded, set this parameter to a large value.</li><li>This value can contain a capacity unit. For example, <strong id="b20296417273"><a name="b20296417273"></a><a name="b20296417273"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p177475153515"><a name="p177475153515"></a><a name="p177475153515"></a>The recommended value ranges from 4096 to 65536.</p>
</td>
</tr>
<tr id="row139031841164310"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1290354104316"><a name="p1290354104316"></a><a name="p1290354104316"></a>recordMaxLogSize</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p354061110449"><a name="p354061110449"></a><a name="p354061110449"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p173721511451"><a name="p173721511451"></a><a name="p173721511451"></a>Size of a result list containing success, failure, or warning lists in a batch task, in bytes. The default value is 30 MB.</p>
<div class="note" id="note206161830145111"><a name="note206161830145111"></a><a name="note206161830145111"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p9617113045113"><a name="p9617113045113"></a><a name="p9617113045113"></a>This value can contain a capacity unit. For example, <strong id="b01321527579"><a name="b01321527579"></a><a name="b01321527579"></a>1 MB</strong> indicates 1048576 bytes.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p15903241164314"><a name="p15903241164314"></a><a name="p15903241164314"></a>The recommended value ranges from 5 MB to 100 MB.</p>
</td>
</tr>
<tr id="row12905004317"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p4295508439"><a name="p4295508439"></a><a name="p4295508439"></a>recordBackups</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1549141116442"><a name="p1549141116442"></a><a name="p1549141116442"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p8465154124517"><a name="p8465154124517"></a><a name="p8465154124517"></a>Maximum number of result lists of successful or failed batch tasks that can be retained. The default value is 1024.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p182945034313"><a name="p182945034313"></a><a name="p182945034313"></a>N/A</p>
</td>
</tr>
<tr id="row027565644416"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1927695624419"><a name="p1927695624419"></a><a name="p1927695624419"></a>humanReadableFormat</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p15932114458"><a name="p15932114458"></a><a name="p15932114458"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p14276956164412"><a name="p14276956164412"></a><a name="p14276956164412"></a>Indicates whether to convert the number of bytes in the object listing result and result list content to the human-readable format. The default value is <strong id="b138371916181811"><a name="b138371916181811"></a><a name="b138371916181811"></a>true</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p627635634414"><a name="p627635634414"></a><a name="p627635634414"></a>N/A</p>
</td>
</tr>
<tr id="row9289115614439"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1129025634314"><a name="p1129025634314"></a><a name="p1129025634314"></a>showProgressBar</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1129095654315"><a name="p1129095654315"></a><a name="p1129095654315"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p82901256134314"><a name="p82901256134314"></a><a name="p82901256134314"></a>Indicates whether to display the progress bar on the console. The value <strong id="b2336149163714"><a name="b2336149163714"></a><a name="b2336149163714"></a>true</strong> indicates that the progress bar is displayed. The default value is <strong id="b16349189113716"><a name="b16349189113716"></a><a name="b16349189113716"></a>true</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p1629095624312"><a name="p1629095624312"></a><a name="p1629095624312"></a>N/A</p>
</td>
</tr>
<tr id="row3984101315382"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p1998421323818"><a name="p1998421323818"></a><a name="p1998421323818"></a>showStartTime</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p298441311387"><a name="p298441311387"></a><a name="p298441311387"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p09841138381"><a name="p09841138381"></a><a name="p09841138381"></a>Indicates whether to display the start time on the console. The value <strong id="b544204313719"><a name="b544204313719"></a><a name="b544204313719"></a>true</strong> indicates that start time is displayed. The default value is <strong id="b931591218381"><a name="b931591218381"></a><a name="b931591218381"></a>true</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p58383208389"><a name="p58383208389"></a><a name="p58383208389"></a>N/A</p>
</td>
</tr>
<tr id="row1534531794316"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p143460172438"><a name="p143460172438"></a><a name="p143460172438"></a>colorfulProgress</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p7346317184311"><a name="p7346317184311"></a><a name="p7346317184311"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p534616174439"><a name="p534616174439"></a><a name="p534616174439"></a>Indicates whether to enable the progress bar with colors. The value <strong id="b19997889358"><a name="b19997889358"></a><a name="b19997889358"></a>true</strong> indicates that the bar is enabled.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p33461317114320"><a name="p33461317114320"></a><a name="p33461317114320"></a>N/A</p>
</td>
</tr>
<tr id="row45151020122210"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p16516142012213"><a name="p16516142012213"></a><a name="p16516142012213"></a>helpLanguage</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p55168208221"><a name="p55168208221"></a><a name="p55168208221"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p85161420112219"><a name="p85161420112219"></a><a name="p85161420112219"></a>Language of the help documents. Options are as follows:</p>
<a name="ul105883424222"></a><a name="ul105883424222"></a><ul id="ul105883424222"><li>Chinese</li><li>English</li></ul>
<p id="p44701718758"><a name="p44701718758"></a><a name="p44701718758"></a>The default value is <strong id="b194381111583"><a name="b194381111583"></a><a name="b194381111583"></a>English</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p36381835122210"><a name="p36381835122210"></a><a name="p36381835122210"></a>N/A</p>
</td>
</tr>
<tr id="row781417083519"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p15815103352"><a name="p15815103352"></a><a name="p15815103352"></a>defaultTempFileDir</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1781513010355"><a name="p1781513010355"></a><a name="p1781513010355"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p14815170203519"><a name="p14815170203519"></a><a name="p14815170203519"></a>Indicates the directory for storing temporary files during download.</p>
<p id="p16698926194511"><a name="p16698926194511"></a><a name="p16698926194511"></a>The default value is the <strong id="b17697141115411"><a name="b17697141115411"></a><a name="b17697141115411"></a>.obsutil_tempfile</strong> subfolder in the user directory (<strong id="b1084318141211"><a name="b1084318141211"></a><a name="b1084318141211"></a>HOME</strong> in Linux or macOS and <strong id="b1020691213215"><a name="b1020691213215"></a><a name="b1020691213215"></a>C:\Users\</strong><em id="i18698122613459"><a name="i18698122613459"></a><a name="i18698122613459"></a>&lt;Username&gt;</em> in Windows).</p>
<div class="note" id="note7556132215312"><a name="note7556132215312"></a><a name="note7556132215312"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul15994748142619"></a><a name="ul15994748142619"></a><ul id="ul15994748142619"><li>Temporary files generated during multipart download are stored in this directory. Therefore, ensure that the user who executes obsutil has the write permission on the path.</li><li>The available space of the partition where the path is located must be greater than the size of the objects to be downloaded.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p58157093511"><a name="p58157093511"></a><a name="p58157093511"></a>N/A</p>
</td>
</tr>
<tr id="row18422132184015"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p8423152104019"><a name="p8423152104019"></a><a name="p8423152104019"></a>checkSourceChange</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1012131194013"><a name="p1012131194013"></a><a name="p1012131194013"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p194231423407"><a name="p194231423407"></a><a name="p194231423407"></a>Indicates whether to check the change of source files or objects during upload/download/copy. The value <strong id="b2061652151219"><a name="b2061652151219"></a><a name="b2061652151219"></a>true</strong> indicates that the function is enabled.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p04239204012"><a name="p04239204012"></a><a name="p04239204012"></a>N/A</p>
</td>
</tr>
<tr id="row87186322412"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p147197321047"><a name="p147197321047"></a><a name="p147197321047"></a>skipCheckEmptyFolder</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p15719532149"><a name="p15719532149"></a><a name="p15719532149"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p187191632642"><a name="p187191632642"></a><a name="p187191632642"></a>Indicates whether to skip checking empty folders on the OBS server during download. <strong id="b126992071966"><a name="b126992071966"></a><a name="b126992071966"></a>true</strong> indicates to skip the check. The default value is <strong id="b186414361267"><a name="b186414361267"></a><a name="b186414361267"></a>false</strong>.</p>
<div class="notice" id="note14592736086"><a name="note14592736086"></a><a name="note14592736086"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p171891211988"><a name="p171891211988"></a><a name="p171891211988"></a>If this parameter is set to <strong id="b1421215541063"><a name="b1421215541063"></a><a name="b1421215541063"></a>true</strong>, the directory structure downloaded to your local PC may be different from that in OBS.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p13719532441"><a name="p13719532441"></a><a name="p13719532441"></a>N/A</p>
</td>
</tr>
<tr id="row4298124583"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p429984184"><a name="p429984184"></a><a name="p429984184"></a>fsyncForDownload</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p8299741783"><a name="p8299741783"></a><a name="p8299741783"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p15299140820"><a name="p15299140820"></a><a name="p15299140820"></a>Indicates whether to forcibly synchronize memory data to disks during download. The value <strong id="b17439141112316"><a name="b17439141112316"></a><a name="b17439141112316"></a>true</strong> indicates to enable forcible synchronization. The default value is <strong id="b19391911132418"><a name="b19391911132418"></a><a name="b19391911132418"></a>false</strong>.</p>
<div class="note" id="note1043555818126"><a name="note1043555818126"></a><a name="note1043555818126"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul1830874141311"></a><a name="ul1830874141311"></a><ul id="ul1830874141311"><li>Set this parameter to <strong id="b1621561262"><a name="b1621561262"></a><a name="b1621561262"></a>true</strong> for scenarios that require high data reliability.</li><li>If this parameter is set to <strong id="b19600713152711"><a name="b19600713152711"></a><a name="b19600713152711"></a>true</strong>, the download performance will be deteriorated. Therefore, exercise caution when using this parameter.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p12657181217810"><a name="p12657181217810"></a><a name="p12657181217810"></a>N/A</p>
</td>
</tr>
<tr id="row97408551683"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p7740855382"><a name="p7740855382"></a><a name="p7740855382"></a>memoryEconomicalScanForUpload</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p974165515820"><a name="p974165515820"></a><a name="p974165515820"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p18741955287"><a name="p18741955287"></a><a name="p18741955287"></a>Indicates whether to use the scanning mode that occupies less memory space when uploading a folder. The value <strong id="b1262185619394"><a name="b1262185619394"></a><a name="b1262185619394"></a>true</strong> indicates using this method and the default value <strong id="b28321815408"><a name="b28321815408"></a><a name="b28321815408"></a>false</strong> indicates not using this method.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p67417551815"><a name="p67417551815"></a><a name="p67417551815"></a>N/A</p>
</td>
</tr>
<tr id="row854835710124"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p15549165741211"><a name="p15549165741211"></a><a name="p15549165741211"></a>forceOverwriteForDownload</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p65498575122"><a name="p65498575122"></a><a name="p65498575122"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p135499571127"><a name="p135499571127"></a><a name="p135499571127"></a>Indicates to forcibly overwrite the local executable file (even if the local executable file is running) when downloading objects to the Linux OS or macOS. The value <strong id="b10652654142116"><a name="b10652654142116"></a><a name="b10652654142116"></a>true</strong> means to overwrite, and the default value is <strong id="b135511916142216"><a name="b135511916142216"></a><a name="b135511916142216"></a>true</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p991816196133"><a name="p991816196133"></a><a name="p991816196133"></a>N/A</p>
</td>
</tr>
<tr id="row122076369221"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p520833642219"><a name="p520833642219"></a><a name="p520833642219"></a>panicForSymbolicLinkCircle</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p1920823642220"><a name="p1920823642220"></a><a name="p1920823642220"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p12208153615224"><a name="p12208153615224"></a><a name="p12208153615224"></a>Indicates the processing method after a symbolic link loop is detected during upload. The value <strong id="b155922810100"><a name="b155922810100"></a><a name="b155922810100"></a>false</strong> indicates that errors are only recorded. The value <strong id="b043514262105"><a name="b043514262105"></a><a name="b043514262105"></a>true</strong> indicates that panic is triggered. The default value is <strong id="b189368520103"><a name="b189368520103"></a><a name="b189368520103"></a>false</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p6892203217258"><a name="p6892203217258"></a><a name="p6892203217258"></a>N/A</p>
</td>
</tr>
<tr id="row1684111384236"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p88411738152317"><a name="p88411738152317"></a><a name="p88411738152317"></a>fastFailThreshold</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p187111822182618"><a name="p187111822182618"></a><a name="p187111822182618"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p16841203810233"><a name="p16841203810233"></a><a name="p16841203810233"></a>Threshold for fast failure upon 4XX errors of batch tasks. When the number of 4XX errors exceeds the threshold, the fast failure process is triggered. All tasks that are not executed or being scanned are suspended. The default value is <strong id="b777818361125"><a name="b777818361125"></a><a name="b777818361125"></a>5</strong>.</p>
<div class="note" id="note13911725264"><a name="note13911725264"></a><a name="note13911725264"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p057045503518"><a name="p057045503518"></a><a name="p057045503518"></a>The fast failure mechanism is to avoid excessive traffic generated during batch task execution. To start a fast failure as soon as possible, set this parameter to <strong id="b125291340169"><a name="b125291340169"></a><a name="b125291340169"></a>0</strong> or <strong id="b2947939161613"><a name="b2947939161613"></a><a name="b2947939161613"></a>-1</strong>, indicating that the fast failure process starts immediately whenever a 4XX error occurs.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p19841133818230"><a name="p19841133818230"></a><a name="p19841133818230"></a>N/A</p>
</td>
</tr>
<tr id="row14330834194012"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p20331834204016"><a name="p20331834204016"></a><a name="p20331834204016"></a>abortHttpStatusForResumableTasks</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p833113349401"><a name="p833113349401"></a><a name="p833113349401"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p933116346400"><a name="p933116346400"></a><a name="p933116346400"></a>HTTP status codes for fast interruption of multipart upload, download, and copy tasks. If a sub-task of a multipart task receives an HTTP code that falls into this range, the multipart task is immediately interrupted. The default values are 401, 403, 404, 405, and 409.</p>
<div class="note" id="note103446484719"><a name="note103446484719"></a><a name="note103446484719"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul2065412944719"></a><a name="ul2065412944719"></a><ul id="ul2065412944719"><li>Multiple HTTP status codes can be carried and separated by commas (,). For example: 401,403,404;</li><li>The status code must be a 4XX HTTP status code. Other status codes are ignored.</li></ul>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p133311434174010"><a name="p133311434174010"></a><a name="p133311434174010"></a>Default value</p>
</td>
</tr>
<tr id="row819881118240"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p61991011142412"><a name="p61991011142412"></a><a name="p61991011142412"></a>showBytesForCopy</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p91996113246"><a name="p91996113246"></a><a name="p91996113246"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p131991711182412"><a name="p131991711182412"></a><a name="p131991711182412"></a>Indicates whether the progress bar displays the rate in bytes when objects are copied between buckets. The default value is <strong id="b140619205228"><a name="b140619205228"></a><a name="b140619205228"></a>false</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p96607311247"><a name="p96607311247"></a><a name="p96607311247"></a>N/A</p>
</td>
</tr>
<tr id="row218965617417"><td class="cellrowborder" valign="top" width="15%" headers="mcps1.2.5.1.1 "><p id="p12190155616416"><a name="p12190155616416"></a><a name="p12190155616416"></a>proxyUrl</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.2.5.1.2 "><p id="p101903561542"><a name="p101903561542"></a><a name="p101903561542"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="42%" headers="mcps1.2.5.1.3 "><p id="p419014561419"><a name="p419014561419"></a><a name="p419014561419"></a>HTTP proxy example: <strong id="b6646174334012"><a name="b6646174334012"></a><a name="b6646174334012"></a>http://username:password@your-proxy:8080</strong></p>
<div class="note" id="note10614125201319"><a name="note10614125201319"></a><a name="note10614125201319"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p961518529135"><a name="p961518529135"></a><a name="p961518529135"></a>The user name and password cannot contain colons (:) and at signs (@), which will result in parsing errors.</p>
</div></div>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.2.5.1.4 "><p id="p0191556544"><a name="p0191556544"></a><a name="p0191556544"></a>N/A</p>
</td>
</tr>
</tbody>
</table>

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   Set parameters with  **N/A**  as the recommended value based on your needs.  
>-   You are advised to specify  **sdkLogPath**  and  **utilLogPath**  to enable SDK logging and obsutil logging.  
>-   The values of  **defaultBigfileThreshold**,  **defaultPartSize**,  **rateLimitThreshold**,  **sdkMaxLogSize**,  **utilMaxLogSize**,  **recordMaxLogSize**,  **readBufferIoSize**, and  **writeBufferIoSizecan**  can contain a capacity unit. For example, 1 MB indicates 1048576 bytes.  

