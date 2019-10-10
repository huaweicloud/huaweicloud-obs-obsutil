# Archiving Log Files<a name="EN-US_TOPIC_0149246214"></a>

## Function<a name="section1479112110815"></a>

You can use this command to archive log files to a local PC or to a specified bucket.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Archiving to a local PC

        ```
        obsutil archive [file_or_folder_url] [-config=xxx]
        ```

    -   Archiving to a specified bucket

        ```
        obsutil archive obs://bucket[/key] [-config=xxx]
        ```


-   In Linux or macOS
    -   Archiving to a local PC

        ```
        obsutil archive [file_or_folder_url] [-config=xxx]
        ```

    -   Archiving to a specified bucket

        ```
        obsutil archive obs://bucket[/key] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="25%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="59%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>file_or_folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="59%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Indicates the path to which log files are archived. The rules are as follows:</p>
<a name="ul105928456532"></a><a name="ul105928456532"></a><ul id="ul105928456532"><li>If this parameter is left blank, log files are archived to the same directory where obsutil commands reside with <strong id="b1549218199916"><a name="b1549218199916"></a><a name="b1549218199916"></a>obsutil_log.zip</strong> as the archive file name.</li><li>If this parameter specifies a file or folder path that does not exist, the tool checks whether the value ends with a slash (/) or backslash (\). If yes, a folder is created based on the path, and log files are archived to the newly created directory with <strong id="b19182195312314"><a name="b19182195312314"></a><a name="b19182195312314"></a>obsutil_log.zip</strong> as the archive file name.</li><li>If this parameter specifies a file or folder path that does not exist and the value does not end with a slash (/) or backslash (\), log files are archived to a local PC with the value as the archive file name.</li><li>If this parameter specifies an existing .zip file, then log files are archived to a local PC overwriting the existing file, with the value as the archive file name.</li><li>If this parameter specifies an existing folder, then log files are archived to the specified directory with <strong id="b94473916314"><a name="b94473916314"></a><a name="b94473916314"></a>obsutil_log.zip</strong> as the archive file name.</li></ul>
<div class="note" id="note2947114643"><a name="note2947114643"></a><a name="note2947114643"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p394734546"><a name="p394734546"></a><a name="p394734546"></a>All archive files are .zip files.</p>
</div></div>
</td>
</tr>
<tr id="row1583413141821"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p154092113211"><a name="p154092113211"></a><a name="p154092113211"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p105401120328"><a name="p105401120328"></a><a name="p105401120328"></a>Mandatory for archiving log files to a specified bucket</p>
</td>
<td class="cellrowborder" valign="top" width="59%" headers="mcps1.1.4.1.3 "><p id="p1954062113212"><a name="p1954062113212"></a><a name="p1954062113212"></a>Bucket name</p>
</td>
</tr>
<tr id="row941361618217"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p13233171113215"><a name="p13233171113215"></a><a name="p13233171113215"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p152331511173216"><a name="p152331511173216"></a><a name="p152331511173216"></a>Optional for archiving log files to a specified bucket</p>
</td>
<td class="cellrowborder" valign="top" width="59%" headers="mcps1.1.4.1.3 "><p id="p16965474526"><a name="p16965474526"></a><a name="p16965474526"></a>Indicates the object name or object name prefix when archiving log files to a specified bucket.</p>
<p id="p060018221533"><a name="p060018221533"></a><a name="p060018221533"></a>The rules are as follows:</p>
<a name="ul7190122515538"></a><a name="ul7190122515538"></a><ul id="ul7190122515538"><li>If this parameter is left blank, log files are archived to the root directory of the bucket with <strong id="b4475192719426"><a name="b4475192719426"></a><a name="b4475192719426"></a>obsutil_log.zip</strong> as the object name.</li><li>If the value ends with a slash (/), the value is used as the object name prefix when archiving log files, and the object name is the value plus <strong id="b1248317223469"><a name="b1248317223469"></a><a name="b1248317223469"></a>obsutil_log.zip</strong>. Otherwise, log files are archived with the value as the object name.</li></ul>
</td>
</tr>
<tr id="row3739132910383"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="59%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil archive**  command to archive log files to the same directory where the tool is executed.

    ```
    obsutil archive
    
    [----------------------------------------------------------] 100.00% 15/15 35ms
    Succeed to archive log files to [D:\obsutil\obsutil_log.zip]
    ```


