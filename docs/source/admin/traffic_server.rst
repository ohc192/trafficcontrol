..
..
.. Licensed under the Apache License, Version 2.0 (the "License");
.. you may not use this file except in compliance with the License.
.. You may obtain a copy of the License at
..
..     http://www.apache.org/licenses/LICENSE-2.0
..
.. Unless required by applicable law or agreed to in writing, software
.. distributed under the License is distributed on an "AS IS" BASIS,
.. WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
.. See the License for the specific language governing permissions and
.. limitations under the License.
..

*****************************
Traffic Server Administration
*****************************
Installing Traffic Server
=========================

#. Build the Traffic Server RPM. The best way to do this is to follow the Traffic Server documents: ::

    https://docs.trafficserver.apache.org/en/latest/getting-started/index.en.html#installation

#. Build the astats RPM using the appropriate version number: ::

    https://github.com/apache/trafficcontrol/tree/<version>/traffic_server

   Sample link: ::

     https://github.com/apache/trafficcontrol/tree/master/traffic_server

#.  Install Traffic Server and astats: ::

        sudo yum -y install trafficserver-*.rpm astats_over_http*.rpm

#.  Add the server using the Traffic Ops web interface:

    #. Select **Servers**.
    #. Scroll to the bottom of the page and click **Add Server**.
    #. Complete the "Required Info:" section:
        * Set 'Interface Name' to the name of the interface from which traffic server delivers content.
        * Set 'Type' to 'MID' or 'EDGE'.
    #. Click **Submit**.
    #. Click **Save**.
    #. Click **Online Server**.
    #. Verify that the server status is now listed as **Reported**

#.  Install the ORT script and run it in 'badass' mode to create the initial configuration, see :ref:`reference-traffic-ops-ort`

#.  Start the service: ``sudo service trafficserver start``

#.  Configure traffic server to start automatically: ``sudo systemctl enable trafficserver``

#.  Verify that the installation is good:

    #. Make sure that the service is running: ``sudo systemctl status trafficserver``

    #. Assuming a traffic monitor is already installed, browse to it, i.e. http://<trafficmonitorURL>, and verify that the traffic server appears in the "Cache States" table, in white.


.. _reference-traffic-ops-ort:

Configuring Traffic Server
==========================
All of the Traffic Server application configuration files are generated by Traffic Ops and installed by way of the traffic_ops_ort.pl script.
The traffic_ops_ort.pl should be installed on all caches (by puppet or other non Traffic Ops means), usually in /opt/ort. It is used to do the initial install of the config files when the cache is being deployed, and to keep the config files up to date when the cache is already in service.  The usage message of the script is shown below: ::

    $ sudo /opt/ort/traffic_ops_ort.pl
    ====-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-====
    Usage: ./traffic_ops_ort.pl <Mode> <Log_Level> <Traffic_Ops_URL> <Traffic_Ops_Login> [optional flags]
      <Mode> = interactive - asks questions during config process.
      <Mode> = report - prints config differences and exits.
      <Mode> = badass - attempts to fix all config differences that it can.
      <Mode> = syncds - syncs delivery services with what is configured in Traffic Ops.
      <Mode> = revalidate - checks for updated revalidations in Traffic Ops and applies them.  Requires Traffic Ops 2.1.

      <Log_Level> => ALL, TRACE, DEBUG, INFO, WARN, ERROR, FATAL, NONE

      <Traffic_Ops_URL> = URL to Traffic Ops host. Example: https://trafficops.company.net

      <Traffic_Ops_Login> => Example: 'username:password'

      [optional flags]:
        dispersion=<time>      => wait a random number between 0 and <time> before starting. Default = 300.
        login_dispersion=<time>  => wait a random number between 0 and <time> before login. Default = 0.
        retries=<number>       => retry connection to Traffic Ops URL <number> times. Default = 3.
        wait_for_parents=<0|1> => do not update if parent_pending = 1 in the update json. Default = 1, wait for parents.
    ====-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-====
    $

Installing the ORT script
--------------------------

#.  Build the ORT script RPM from the Apache Build Server and install it: ::

        https://builds.apache.org/view/S-Z/view/TrafficControl/

    Sample command (NOTE that the example name used here most likely no longer exists): ::

        sudo wget https://builds.apache.org/job/trafficcontrol-2.1.x-build/lastSuccessfulBuild/artifact/dist/traffic_ops_ort-2.1.0-6807.1dcd512f.el7.x86_64.rpm
        sudo yum install traffic_ops_ort*.rpm

#.  Install modules required by ORT if needed: ``sudo yum -y install perl-JSON perl-Crypt-SSLeay``

#.  For initial configuration or when major changes (like a Profile change) need to be made, run the script in "badass mode". All required rpm packages
    will be installed, all Traffic Server config files will be fetched and installed, and (if needed) the Traffic Server application will be restarted.

    Example run below: ::

        $ sudo /opt/ort/traffic_ops_ort.pl --dispersion=0 badass warn https://ops.$tcDomain admin:admin123

    .. Note:: First run gives a lot of state errors that are expected. The "badass" mode fixes these issue s. Run it a second time, this should be cleaner.
       Also, note that many ERROR messages emitted by ORT are actually information messages. Do not panic.

#.  Create a cron entry for running ort in 'syncds' mode every 15 minutes.
    This makes traffic control check periodically if 'Queue Updates' was run on Traffic Ops, and it so, get the updated configuration.

    Run ``sudo crontab -e`` and add the following line ::

  	  	*/15 * * * * /opt/ort/traffic_ops_ort.pl syncds warn https://traffops.kabletown.net admin:password --login_dispersion=30 --dispersion=180 > /tmp/ort/syncds.log 2>&1

    Changing ``https://traffops.kabletown.net``, ``admin``, and ``password`` to your CDN URL and credentials.

    .. Note:: By default, running ort on an edge traffic server waits for it's parent (mid) servers to download their configuration before
       it downloads it's own configuration. Because of this, scheduling ort for running every 15 minutes (with 5 minutes default dispersion) means
       that it might take up to ~35 minutes for a "Queue Updates" operation to affect all traffic servers.  To customize this dispersion time, use
       the command line option --dispersion=x where x is the number of seconds for the dispersion period.  Servers will select a random number from
       within this dispersion period to being pulling down configuration files from Traffic Ops.  Another option, --login_dispersion=x can be used.
       This option creates a dispersion period after the job begins during which ORT will wait before logging in and checking Traffic Ops for updates
       to the server.  This defaults to 0.  If use_reval_pending, a.k.a. Rapid Revalidate is enabled, edges will NOT wait for their parents to download
       their configuration before downloading their own.

    .. Note:: In 'syncds' mode, the ort script updates only configurations that might be changed as part of normal operations, such as:

        * Delivery Services
        * SSL certificates
        * Traffic Monitor IP addresses
        * Logging configuration
        * Revalidation requests (By default. If Rapid Revalidate is enabled, this will only be checked by using a separate revalidate command in ORT.)


#.  If Rapid Revalidate is enabled in Traffic Ops, create a second cron job for revalidation checks.  ORT will not check revalidation files if Rapid Revalidate
    is enabled. This setting allows for a separate check to be performed every 60 seconds to verify if a revalidation update has been made.

    Run ``sudo crontab -e`` and add the following line ::

        */1 * * * * /opt/ort/traffic_ops_ort.pl revalidate warn https://traffops.kabletown.net admin:password --login_dispersion=30 > /tmp/ort/syncds.log 2>&1
