ETransV2
==============

.. image:: https://goreportcard.com/badge/github.com/Class04OpenSourceORG/ETransV2
   :target: https://goreportcard.com/report/github.com/Class04OpenSourceORG/ETransV2

ETransV2 is a data transmit tools written in Go.

.. image:: https://github.com/Class04OpenSourceORG/ETransV2/blob/master/t.png

Install
--------------

**For Server** ::

  git clone https://github.com/Class04OpenSourceORG/ETransV2
  cd ETransV2
  cd server
  go install

**For Client** ::

  git clone https://github.com/Class04OpenSourceORG/ETransV2
  cd ETransV2
  cd client
  go install
  
Usage
--------------

**Server**

``server <listening_port>``

**Client**

``client <server-ip:server-port> <filepath>``

BenchMark
--------------

Transport a 700MB file in LAN.

* Average Speed: **77Mbps**
* Top Speed: **102Mbps**
* Lowest Speed: **67Mbps**
* RAM Use: **2.4MB** (Both Client & Server)

See Also
--------------

.. _ETransV3: https://github.com/Class04OpenSourceORG
