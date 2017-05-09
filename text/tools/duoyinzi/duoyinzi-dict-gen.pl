#! /usr/bin/env perl
# -*- perl -*- 

# @Author: Tong SUN, (c)2017, all right reserved

# ============================================================== &us ===
# ............................................................. Uses ...

# -- global directives/modules
use strict;			# !
use warnings;

# ============================================================== &sb ===
# .................................................... Script begins ...

my $fh = \*DATA;

while(<$fh>) {
    next if /^\s*#/;
    next if /^\s*$/;
    my @F = split(//);
    chomp;
    print $_. ": ". join(" ", @F);
}

# == end of program

__DATA__
ABCDG
ABCDGM
ABCDK
ABCDKUV
ABCDKXY
ABCF
ABCGHIJK
ABCGLL
AK
ALMNOPQ
AMOQ
BDD
EFGHIJ
