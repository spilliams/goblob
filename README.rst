goblob
======

A little tool to turn SQL blob strings into JSON

Install it with

.. code-block:: shell

   go get -u github.com/spilliams/goblob


and then use it either by passing args

.. code-block:: shell

   $ goblob 'a:1:{s:3:"foo";s:3:"bar";}'
   {
     "foo":"bar"
   }

or with a pipe


.. code-block:: shell

   $ echo 'a:1:{s:3:"foo";s:3:"bar";}' | goblob
   {
     "foo":"bar"
   }
   
   $ echo 'a:1:{s:3:"baz";s:3:"qux";}' >> file
   $ cat file | goblob
   {
     "baz":"qux"
   }

(and it's ready for further piping, for instance into `jq <https://stedolan.github.io/jq/>`_)
