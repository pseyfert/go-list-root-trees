# go-list-root-trees

[![Licence: GPL v3](https://img.shields.io/github/license/pseyfert/go-list-root-trees.svg)](LICENSE)
[![travis Status](https://travis-ci.org/pseyfert/go-list-root-trees.svg?branch=master)](https://travis-ci.org/pseyfert/go-list-root-trees)

This is a command line tool to list all trees in a root file.

The behaviour should be similar to that of `root_numpy.list_trees`. That is,
trees in subdirectories in a file are listed.

This should shorten a user's lookups traversing subdirectories.

The tool is intended to ease usage of other command line programs that read
trees from root files, such as DooSelection tools or the
[tmva-branch-adder](https://github.com/pseyfert/tmva-branch-adder). It can be
used for tab completion functions.

printout looks like:
```
pseyfert@robusta ~/coding/tmva-branch-adder > root-ls-tree /afs/cern.ch/user/p/pseyfert/DTT.root
EventTuple/EventTuple	B02DD/DecayTree	GetIntegratedLuminosity/LumiTuple
```

# TODO:

 - for the intended use case of tab completions, low latency is key, so a
   closer look at performance should be taken at some point.
 - possibly an argument to prepend path names to the printout in a root
   standard format (`rootls` accepts `<filename>.root:<dir>/<tree>`) So for
   some interfaces having filename and tree path together might be more
   convenient:
```
pseyfert@robusta ~/coding/tmva-branch-adder > root-ls-tree -f /afs/cern.ch/user/p/pseyfert/DTT.root
/afs/cern.ch/user/p/pseyfert/DTT.root:EventTuple/EventTuple
/afs/cern.ch/user/p/pseyfert/DTT.root:B02DD/DecayTree
/afs/cern.ch/user/p/pseyfert/DTT.root:GetIntegratedLuminosity/LumiTuple
```
