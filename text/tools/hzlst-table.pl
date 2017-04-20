#! /usr/bin/env perl
# -*- perl -*- 

# {{{ LICENSE: 

# 
# Permission to use, copy, modify, and distribute this software and its
# documentation for any purpose and without fee is hereby granted, provided
# that the above copyright notices appear in all copies and that both those
# copyright notices and this permission notice appear in supporting
# documentation, and that the names of author not be used in advertising or
# publicity pertaining to distribution of the software without specific,
# written prior permission.  Tong Sun makes no representations about the
# suitability of this software for any purpose.  It is provided "as is"
# without express or implied warranty.
#
# TONG SUN DISCLAIM ALL WARRANTIES WITH REGARD TO THIS SOFTWARE, INCLUDING ALL
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS, IN NO EVENT SHALL ADOBE
# SYSTEMS INCORPORATED AND DIGITAL EQUIPMENT CORPORATION BE LIABLE FOR ANY
# SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER
# RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF
# CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN
# CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
# 

# }}} 


=head1 Introduction
  
  This program list all GB2312 encoded HZ characters
  
=cut

$cpl=16;        # chars per line
@qm = ( 0xa1..0xfe );
@wm = ( 0xa1..0xfe );

# print header ruler
foreach $ndx ( (1..$cpl) ){
        printf "%02s ", $ndx;
        }
print "\n";
foreach $ndx ( (1..$cpl) ){
        printf "%02x ", $ndx;
        }
print "\n\n";

# print contents
foreach $qm ( @qm ){
    $ndx=0;
    printf "%d (0x%x) \n", $qm,$qm;
    foreach $wm (@wm) {
        print chr($qm) .chr($wm). " ";
                if (++$ndx >=$cpl ){
            print "\n";
            $ndx=0;}
        }
    print "\n\n";
    }
    
