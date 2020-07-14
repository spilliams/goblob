goblob
======

A little tool to turn SQL blob strings into JSON

Install it with

.. code-block:: console
   go get -u github.com/spilliams/goblob


and then use it either by passing args

.. code-block:: console
   $ goblob 'a:1:{s:3:"foo";s:3:"bar";}'
   {"foo":"bar"}

or with a pipe


.. code-block:: console
   $ echo 'a:1:{s:3:"foo";s:3:"bar";}' | goblob
   {"foo":"bar"}
   
   $ echo 'a:1:{s:3:"baz";s:3:"qux";}' >> file
   $ cat file | goblob
   {"baz":"qux"}
