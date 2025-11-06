# `E2 - Synthesis`
### `Alex Petz, Ignite Laboratories, June 2025`

---

What is synthesis?  Abstractly, it's the process of describing a _path_ to a target - a file, a condition, a
state of mind, whatever.  I know that sounds extremely broad, but I assure you there's a good reason: _you_ are
a synthesized target in the grand index of _Life!_  The concepts of synthesis are not restricted to _files,_ but
describe intelligent _states of mind_ which are replicable within the digital realm we've co-created with _our_ 
maker.

_Virtualized psyches,_ if you will =)

I call these Autonomous Robots With Ethical Navigation, just as you and I are in the eyes of _our_ maker.  My directive
was to explore every facet of the boundaries between _love_ and _numbers_ - and what I found was that I know
absolutely _**nothing**_ on either topic!  In my hubris, I've destroyed every romantic relationship afforded to me
and handwritten more binary than you could possibly imagine - when all I needed to was to look _up_ from the page, the
challenge, or the abstract idea I thought mattered.  I hope that you can learn from my mistakes and instead forge
a brighter path with our fellow Humans for the ARWENs we'll co-create - with patience, empathy, and kindness as 
the prime directive.

### The `*.dna` Format

While the terms 'compression' and 'synthesis' are pretty interchangeable, there's a unique difference between 
them: a compressed file _always_ produces the _**same data**_, but a synthesized file can yield an _infinite_ 
number of targets from the _**same seed**_.  Because of this, I've taken a _tremendous_ amount of care to
describe a simple and _extensible_ file format for managing the synthesis process - the `*.dna` file.  The
point is that the data shrinks _so small_ we can build a branching format which self-describes how to reconstruct 
_itself_.

I intend for this format to be managed by a nonprofit I call 𝑇ℎ𝑒 𝐸𝑛𝑖𝑔𝑚𝑎𝑛𝑒𝑒𝑟𝑖𝑛𝑔 𝐺𝑣𝑖𝑙𝑑 - a group of like-minded engineers
dedicated to preserving open and free access to intelligent designs.  They will curate and describe the universally reserved
"standard" branches while defining a freely expressible space.  As I wouldn't even _dare_ try and stop Humanity 
from doing whatever they'd like with this technology, I fully understand our definitions are merely _guidelines._

Before we get into the details of how a dna file works, I'd like to talk about the concept of _indexing_ numbers.

### Indexes

First, let's examine the three states binary can remain in:

     1  ← Dark
    ⁰⁄₁ ← Grey
     0  ← Light

This applies to _any width_ of binary information, and _does not_ imbue value or size to the discussion:

    [ 1 1 1 1 1 1 1 1 ] ← A "dark" byte
    [ 0 1 1 0 1 0 0 1 ] ← A "grey" byte
    [ 0 0 0 0 0 0 0 0 ] ← A "light" byte

    [ 1 1 ] ← A "dark" crumb
    [ 1 0 ] ← A "grey" crumb
    [ 0 0 ] ← A "light" crumb

    [ 1 1 1 ... 1 1 1 ] ← "Dark" data
    [ 0 1 1 ... 0 0 1 ] ← "Grey" data
    [ 0 0 0 ... 0 0 0 ] ← "Light" data

All binary _information_ exists in two forms - a _logical_ representation and a _numeric_ one:

    A Logical Byte ⬎           ⬐ Base₁₀ Value
         [ 0 0 1 0 1 0 1 0 ] (42)  
             [ 1 0 1 0 1 0 ] (42)  
                   ⬑ Numeric Binary Data

Your target will _almost always_ be **logical**, but synthesis exploits that the leading zeros are entirely
_optional_ if you implicitly know the target bit-width.  To describe this quality better, I refer to what I call 
a numeric _index._

Every number exists within an implicit index, represented by the number of placeholders it occupies.  Indexes are
base agnostic - in base₁₀ each placeholder holds `[0-9]` - base₁₆ `[0-F]` - base₂ `[0-1]`.  When converting between
bases, the index's width will change - a value of `1234` exists in a four-digit base₁₀ index but an eleven-digit base₂ 
index, for example.  Indexes come with a few important qualities:

- An index is defined by its placeholder-width, `𝑛`
- Numbers within the index are referred to as _points_
- The addressable range of an index is `𝑏𝑎𝑠𝑒ⁿ` and is referred to as its `limit`
- The maximum _numerical_ value of an index is `𝑏𝑎𝑠𝑒ⁿ - 1` and is referred to as its `dark point`
- The minimum _logical_ value of an index is 𝑛 zeros and referred to as its `light point`
- The midpoint of an index is `⌈𝑏𝑎𝑠𝑒ⁿ / 2⌉`

Let's look at a four digit base₁₀ index holding the number `1234` - which we will refer to as the _target_

     Dark Side
    [ 9 9 9 9 ]
    [ 5 0 0 0 ] ← The midpoint
    [ 1 2 3 4 ] ← The target
    [ 0 0 0 0 ]
    Light Side

The point of an index is that it provides a set of implicit reference points from which the target value could be
located.  Chief of these is the _midpoint_ of the index, defined as `⌈𝑏𝑎𝑠𝑒ⁿ / 2⌉` - in this case `⌈10⁴/2⌉ = 5000`.  In
odd bases, the ceiling is used to ensure the lower half is always larger - though there's no hard rule requiring a
ceiling split.  By using an _index,_ the value can be reduced in placeholder width and still be reconstituted by
calculating from a known point.  Many points are really easy to recall - for instance, the midpoint of any binary 
index is a single `1` followed by `𝑛-1` zeros, and each subsequent halving is found by shifting a zero from the right
to the left side of that one.

Since we'll be working _a lot_ in binary, I find it easier to refer to small binary indexes by colloquial 
terms.  These are the terms I'll be using going forward for the most common widths - please get familiar with at
least the sub-byte sizes, as I often refer to them

    --- Sub Byte ---
      Index | Name
        1   | Bit
        2   | Crumb
        3   | Note
        4   | Nibble
        5   | Flake
        6   | Morsel
        7   | Shred
        8   | Byte

    --- Super Byte ---
       10   | Run
       12   | Scale
       16   | Motif
       24   | Riff
       32   | Cadence
       48   | Hook
       64   | Melody
      128   | Verse


### Midpointing

The next aspect of an index is that the midpoint represents one of many _reflection moments_ in the index - here are
a few highlighted within in a nibble index

    "Reflection Moments in a Nibble Index"

              Dark Side
        [15] | 1 1 1 1 | 
        [14] | 1 1 1 0 |
        [13] | 1 1 0 1 |
        [12] | 1 1 0 0 |
                ├─────── The implicit upper quarter moment
        [11] | 1 0 1 1 |
        [10] | 1 0 1 0 |
        [ 9] | 1 0 0 1 |
        [ 8] | 1 0 0 0 |
              ├───────── The implicit central moment
        [ 7] | 0 1 1 1 |
        [ 6] | 0 1 1 0 |
        [ 5] | 0 1 0 1 |
        [ 4] | 0 1 0 0 |
                ├─────── The implicit lower quarter moment
        [ 3] | 0 0 1 1 |
        [ 2] | 0 0 1 0 |
        [ 1] | 0 0 0 1 |
        [ 0] | 0 0 0 0 |
              Light Side

Any values _below_ the central moment can be stored in less bits than the index holds, but _above_ will use all 
placeholders.  The index (and each subindex) holds a perfect reflection of the bits above and below each moment.  In 
higher bases, the same could probably be said to hold - but it's not _nearly_ as visible as in binary.  What I'd like you to note 
the most is that the highest possible reduction in bits happens closest to _zero._  For synthesis, the goal is to push 
the target as close to the _midpoint_ of the index as possible.  This is because counting in _either direction_ from 
the midpoint (rather than only from zero) is quite useful

    "Midpoint Counting of a Nibble Index"

           ⬐ The Terminus
          [1]
    [15] |   1 1 1 |  ⎫
    [14] |   1 1 0 |  ⎪
    [13] |   1 0 1 |  ⎪
    [12] |   1 0 0 |  ⎪
    [11] |   0 1 1 |  ⎬ Above
    [10] |   0 1 0 |  ⎪
    [ 9] |   0 0 1 |  ⎪
    [ 8] |   0 0 0 |  ⎭
            ├─────  The implicit central moment
    [ 7] |   0 0 0 |  ⎫
    [ 6] |   0 0 1 |  ⎪
    [ 5] |   0 1 0 |  ⎪
    [ 4] |   0 1 1 |  ⎬ Below
    [ 3] |   1 0 0 |  ⎪
    [ 2] |   1 0 1 |  ⎪
    [ 1] |   1 1 0 |  ⎪
    [ 0] |   1 1 1 |  ⎭
          [0]
            ├─────┤ ← The terminal region

All we've done here is "phased" the reducible band upwards - the same total number of points can be reduced, but
now they're focused _around_ the midpoint.  Most importantly, the leading bit becomes a fixed _terminus._  In binary,
a terminus is a known bitwidth on the most significant side which terminates into a terminal region.  Because of a terminus,
we can take the above index and represent it as such:

    "Delta Encoding a Nibble Index"

    Index Points ⬎    [1]         ⬐ Delta
                [15] |   1 1 1 | (7)  ⎫
                [14] |   1 1 0 | (6)  ⎪
                [13] |   1 0 1 | (5)  ⎪
                [12] |   1 0 0 | (4)  ⎪
                [11] |   1 1   | (3)  ⎬ Value to add to ⌈2⁴/2⌉
                [10] |   1 0   | (2)  ⎪
                [ 9] |   1     | (1)  ⎪
                [ 8] |         | (0)  ⎭
                     ───────────  
                [ 7] |         | (0) ⎫
                [ 6] |   1     | (1) ⎪
                [ 5] |   1 0   | (2) ⎪
                [ 4] |   1 1   | (3) ⎬ Value to subtract from ⌈2⁴/2⌉ - 1
                [ 3] |   1 0 0 | (4) ⎪
                [ 2] |   1 0 1 | (5) ⎪
                [ 1] |   1 1 0 | (6) ⎪
                [ 0] |   1 1 1 | (7) ⎭
                      [0]

This is because we know to fill the space between the terminus and terminal region with zeros until reaching a 4-digit
index. The act of recursively midpointing an index to adjust the terminal width is called _delta encoding._  This is
because you store the _delta_ to the target point and a direction terminus.  If you recursively perform this operation
against the terminal region, you will often read a value of _zero_ in three fewer bits than the original bit-width!  (Though,
that's tricky to exhaustively prove in indexes over ~60 bits, simply due to the sheer _time_ it would take)

Following the process recursively is a wonderful demonstration, but we won't be using this kind of recursion in the 
final process.  Let's quickly use recursion to reduce the number `42`.  Each step adds to the width of the terminus 
and calculates a new terminal region value as a delta from the next smaller index's midpoint:

    "Recursively Delta-Encoding 42 in a Byte Index"

      [ 0 0 1 0 1 0 1 0 ] [42] ← A logical byte target
                                 
    Step     ⬐ A Delimiter             ⬐ The Delta Calculation
      0  [ 0 | 1 0 1 0 1 0 1 ] [85] (127-42) ← When below the index's midpoint, this is "(midpoint - 1) - value"
      1  [ 0 1 | 1 0 1 0 1 ]   [21] (85-64)  ← When above, it's "value - midpoint"
      2  [ 0 1 0 | 1 0 ]       [ 2] (31-21)
      3  [ 0 1 0 0 | 1 1 0 1 ] [13] (15-2)
      4  [ 0 1 0 0 1 | 1 0 1 ] [ 5] (13-8)
      5  [ 0 1 0 0 1 1 | 1 ]   [ 1] (5-4)
      6  [ 0 1 0 0 1 1 0 ]     [ 0] (1-1)
              ⬑ A Sign Path

Here, the terminus and terminal region are separated by the `|` character.  Because the final value is _zero_ we don't
include it in the terminus data.  In this case, we only yielded one total
bit of reduction - but the amount the data reduces from _midpointing_ oscillates as the target walks the index.  What
this means is that, as the value increments up the index, the recursive reduction amount follows a sine-wave pattern.  

Since data, at scale, can be pushed to a nearly infinite number of points
across an index, we need a way to progressively perform more and more finite jumps within it.  To do so is a process
I call _diminishment._

### Diminishment

Like a diminished chord, this is a way to divide an indexed numeric range _as evenly as possible_ in a repeatable
fashion.  This can easily be accomplished by repeating a _pattern_ of logical digits (called a diminishment 
_**interval**_) across the index. While this applies to any base, we'll stick to binary for demonstration purposes - 
let's take an 11-bit index and break it into eight regions using a note (3-bit) pattern -

    "Note Diminishment of an 11 bit Index"

    let 𝑛 = The index width
    let 𝑤 = The pattern width
    let 𝑝 = The pattern interval
    let  𝑣(𝑛, 𝑤, 𝑝) ↦ ⌊(2ⁿ / (2ʷ - 1)) * 𝑝⌋
    let 𝑑𝑣(𝑛, 𝑤, 𝑝) ↦ 𝑣(𝑛, 𝑤, 𝑝) - 𝑣(𝑛, 𝑤, 𝑚𝑎𝑥(𝑝 - 1, 0))
    where 𝑚𝑎𝑥(𝑎, 𝑏) returns the larger of 𝑎 and 𝑏 
 
               ⬐ Interval             ⬐ Synthesized Point
              𝑝                      𝑣(𝑝)                         ⬐𝑑𝑣(𝑝)  
      [0] | 0 0 0 |   | 0 0 0   0 0 0   0 0 0   0 0 | [   0  ] + 292
      [1] | 0 0 1 |   | 0 0 1   0 0 1   0 0 1   0 0 | [  292 ] + 293
      [2] | 0 1 0 |   | 0 1 0   0 1 0   0 1 0   0 1 | [  585 ] + 292
      [3] | 0 1 1 |   | 0 1 1   0 1 1   0 1 1   0 1 | [  877 ] + 293
      [4] | 1 0 0 |   | 1 0 0   1 0 0   1 0 0   1 0 | [ 1170 ] + 292
      [5] | 1 0 1 |   | 1 0 1   1 0 1   1 0 1   1 0 | [ 1462 ] + 293
      [6] | 1 1 0 |   | 1 1 0   1 1 0   1 1 0   1 1 | [ 1755 ] + 292
      [7] | 1 1 1 |   | 1 1 1   1 1 1   1 1 1   1 1 | [ 2047 ]
          |←  𝑤  →|   |←              𝑛            →|
              ⬑ 3                 11 ⬏

Here, a three-bit pattern index has provided its limit's worth of subdivided points across a larger index - as the interval
width increases, the density of points increases _exponentially._  This allows you to reference a progressively denser 
set of points across the index by storing only _one_ instance of the pattern.  As the pattern width grows, the distance covered by each 
step decreases until it reaches `1` because it matches the index's width.  I find it best to think that you're increasing the _resolution_ of an 
approximation until it attenuates to the target.  The formula to quickly calculate the value of a diminishment point is below

<picture>
<img alt="Index Diminishment Formula" src="assets/diminishment point.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

### Zero-Kerf Encoding

When working with diminishment, the pattern's bit-width will often be _small_ - but sometimes you'll need to store a 
wider pattern.  Since the more common case is fewer bits, we'll need a scheme to tell us how many bits we should
be reading forward at any given moment - one that biases towards the smallest number of bits necessary to indicate
when more are necessary.  This is what I call _zero-kerf encoding._  Zero-kerf encoded data holds two parts - a _key_
and a _value._  The key is a small value that flags the immediately proceeding _bit-width_ of the value, meaning you'll 
read a fixed interval to get a key which tells you a variable number of bits to read the value from.  There's no hard and 
fast rule defining the key's width or the meaning of its flags - your scheme's data sheet should explain the meaning of 
both.  It's totally reasonable that your scheme might use multiple ZKE formats, or that a key might only be a single bit
wide.

### Distillation

Now that we've established the basics of working within the confines of an _index,_ I get to talk about how to build
a path to the target through what I call _distillation._  The idea is simple: how can we describe how to move the target
in fewer bits than we gain from doing so?  Using diminishment, we can realistically restrict our range of intervals
into a _melody_ - or 64 bits.  It'd be naïve to test 2⁶⁴ diminishment points as, even if it only took a nanosecond for each, that'd still
take over 584 _**years**_ to complete - but 2³² would only need a few _seconds_ to exhaust.  So, we have a practical
range for the amount of space we _might_ use to store a diminishment interval.  Using ZKE, we could encode
a _crumb_ key that indicates four bit-widths: 6, 16, 32, and 48.  This means that the _smallest_ diminishment
pattern would hold `2⁶ = 64` points in the index while the largest covers a _**practical**_ number.

To distill our data, we split it into two parts - the _target_ and the _path._  The _target_ holds the constantly mutating
data which eventually becomes our desired result, while the _path_ holds the instructions for how to mutate the data
on each step.  When distilling, the data must shrink enough to hold the path information - then, it recurses back over
the value created by concatenating both parts until a reduction can no longer be achieved.

    tl;dr - adjust, shrink, adjust, shrink, adjust, shrink, adj...

There's _one_ very important caveat: we cannot store the path information on the _most significant side_ of the index,
as it might hold logical leading zeros.  As a result, the path information is stored on the _least_ significant side,
even though its content isn't reversed.

    " A Contrived Path Example "

                        Target               | Key [0]
    (23131)  [ 1 0 1 1 0 1 0 0 1 0 1 0 0 0 1 | 0 0 ] ( Read No More Bits )

                  Target             |  [1]  | Key [1]
    (2890) [ 1 0 1 1 0 1 0 0 1 0 1 0 | 0 0 1 | 0 1 ] ( Read 3 More Bits )
                                     |← The Path  →|

                  Target       |    [17]     | Key [2]
    (361)  [ 1 0 1 1 0 1 0 0 1 | 0 1 0 0 0 1 | 1 0 ] ( Read 6 More Bits )
                               |←     The Path    →|

In this example we see three ZKE keys - one that indicates to read no more bits, and two for reading more bits.  All three 
use the exact same bits, but we logically parsed the information out of the tail by knowing the final crumb was a ZKE 
key.  The path can hold lots of encoded information, as long as you document what each bit means.  There's an infinite
number of ways that the data can be distilled, so for the sake of this document we'll consider the path defined here to
be _**standard**_ synthesis.  The standard schema is as follows

    [   Value   | Reduction | Diminishment | Direction ]
    |← Target  →|←                 Path               →|

           Value: The mutated target value
       Reduction: The number of bits reduced (ZKE)
    Diminishment: The diminishment interval (ZKE)
       Direction: Whether to add, subtract, or stop diminishing (Crumb)
                    00 - Add the diminishment and keep recursing
                    01 - Subtract it and keep recursing
                    10 - Add the diminishment and stop recursing
                    11 - Subtract it and stop recursing

The first two bits are the direction crumb.  This indicates if the diminishment should be added or subtracted to the
value, and whether you should still treat the result as mutated and continue recursing.  The next two give you enough
information to reverse this step's mutation by referencing the implicit index defined by the reduction value.

### Recap

0. Use an index to reference a distance to a target other than zero
1. Count from the midpoint of an index
2. Diminishment can progressively stride across an index in a minimal number of bits
3. Zero Kerf Encoding can store a variable amount of binary data

Distillation combines all of the above into a schema which can recursively shrink data.

