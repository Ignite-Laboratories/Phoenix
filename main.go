package main

import (
	"fmt"
	"phoenix/core/svc/tiny"

	"git.ignitelabs.net/janos/core/enum/direction/ordinal"
	"git.ignitelabs.net/janos/core/sys/num"
)

func main() {
	tiny.Primitive[int]{}.ShiftLeft(4).Equals()
	tiny.Base[int](2).Multiply()

	tiny.Compare("5", 4)
	tiny.Add(4, "5").Multiply(4)
	tiny.Int.Add(4, "5")
	fmt.Println(Diminishment(44, 0, 1, 1, 0, 1, 0))

	//bytes, _ := os.ReadFile("readme.md")
	//fmt.Println(string(bytes))
	//
	//data := big.NewInt(0).SetBytes(bytes)
	//fmt.Println(data.Text(2))

	// 0 - index the data
	// 1 - build a key signature
}

func Diminishment(width uint, pattern ...num.Bit) num.Measurement {
	return num.NewMeasurementOfPattern(int(width), ordinal.Positive, pattern...)
}

/*
Yes, I was recreating binary - which is quite hilarious, but oh well!  Haha

Now, let's circle back to the tried and tested way of performing synthesis - binary index midpointing!

This process works by identifying the width of the index containing the data initially - the ultimate size of
the data - and then manipulating the target until it's incomprehensibly small.  Along the way, you record the
number of cycles you've performed midpointing and leave behind a self-described binary trail.

The binary trail, however, must be tied to the index size's MOVEMENT!  This is because it's the ONLY implicit
component of calculating the reduction.  So, a 0 indicates the value was BELOW the midpoint - meaning you store
the delta to it.  A 1 indicates it was above.  Every time you perform a midpoint operation, you
are leaving behind a single bit but yielding an index that is AT LEAST one bit smaller.  As you continue this
path, you will (on average) reach ZERO with three bits to spare.

But three bits is NOT enough to encode anything, especially when it's an AVERAGE of three bits!  Instead, the
FINAL step will ALWAYS be whether the result was able to shrink by one bit (0) or two or more bits (1)

Because of this, the binary data will NEVER grow in width because you add a zero when it doesn't shrink - since that's
literally the last value anyways - but it WILL always be able to indicate when the system can SKIP a bit.  Since the average
is, indeed, above ONE - this will much more slowly shrink the binary information than my greedier method before.

A single round will only yield a single bit at most, but subsequent rounds will accumulate a total aggregate of bits
which could hold the number of times the operations occurred - and the zero count to include at the end to rebuild
the final index.

That feels far more intuitive to write than ever before, and it might have something to do with having just mailed
Casha a letter - seriously, I couldn't even LOOK at the screen before, but now I've been compelled to write the logical
reasoning behind non-greedy midpointing.  Now I'm feeling compelled to pull off for a minute again.

I trust God - implicitly.

*/

/*

To prove this, we have to make a few notes that need answered first:

0. If counting from the midpoint, is the delta truly at least 𝑛-1 bits?
1. When midpointing down, does the result truly yield an average of 1 bits? (allowing growth for zeros without issue)
2. Can I show the frequency nature of the midpointing process using a byte?
3. Can diminishment truly be used to push a large binary value into an ideal spot of the frequency band?
4. Is there a theoretical way to define the MINIMUM width that can be distilled?

The things we are trying to accomplish are as such:

0. Using a binary trail created by the midpointing process to distill and produce a new number
1. Using the least significant bit of the binary trail to indicate whether the reduction yielded more than 1 bit
2.


NOTE: The least significant bit would be REMOVED from the next delta on each pass!  It's an ADDED value, meaning we can include it and remove it transiently without affecting the delta relative to the index.

*/
