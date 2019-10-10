# Configuring Scheduled Tasks Using the Crontab Command<a name="EN-US_TOPIC_0145848634"></a>

## Scenario<a name="section146005674214"></a>

Go to the  **/root**  directory at 21:30 every day and upload the  **/src/src1**  folder to bucket  **obs://bucket-test**  in the incremental mode.

## Prerequisites<a name="section16839141103914"></a>

You have properly enabled the scheduled crond service in the Linux OS.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>Run the  **service crond status**  command to check whether the service is enabled.  

## Procedure<a name="section13314111314402"></a>

1.  Run the  **crontab -e**  command to open the configuration file for setting a scheduled task.
2.  Enter the Insert mode to edit the configuration file.

    ```
    30 21 * * * cd /root && nohup ./obsutil cp /src/src1 obs://bucket-test -r -f -u &>obsutil_crond.log &
    ```

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >Assume that the obsutil tool is in the  **/root**  directory. The preceding configuration is described as follows: Go to the  **/root**  directory at 21:30 every day, upload the  **/src/src1**  folder to bucket  **obs://bucket-test**  in incremental mode, and redirect the command output to the  **obsutil\_crond.log**  file in the  **/root**  directory.  

3.  Press  **Esc**  to exit the Insert mode. Then input  **:wq**  and press  **Enter**  to save the configuration and exit.
4.  Run the  **crontab -l**  command to check whether the scheduled task is configured successfully.

## FAQs<a name="section14418131348"></a>

1.  How do I determine whether a scheduled task is being executed?
    -   Run the  **tail /var/log/cron**  command to view the latest scheduled task execution records.
    -   Run the  **ps -ef | grep obsutil**  command to check whether obsutil is being executed.

2.  How do I forcibly stop an ongoing scheduled task?
    1.  Run the  **ps -ef | grep obsutil**  command to check the process of obsutil.
    2.  Run the  **kill -9 **_PID_  command to forcibly stop the process, where  _PID_  indicates the queried process ID.


