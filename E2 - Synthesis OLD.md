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

### The DNA Format

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
base agnostic - in base₁₀ each placeholder holds `[0,9]` - base₁₆ `[0,F]` - base₂ `[0,1]`.  When converting between
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

    "Terminal Counting of a Nibble Index"

    Index Points ⬎    [1]  ⬐ Delta
                [15] |   1 1 1 |  ⎫
                [14] |   1 1 0 |  ⎪
                [13] |   1 0 1 |  ⎪
                [12] |   1 0 0 |  ⎪
                [11] |   1 1   |  ⎬ Value to add to ⌈2⁴/2⌉
                [10] |   1 0   |  ⎪
                [ 9] |   1     |  ⎪
                [ 8] |         |  ⎭
                     ───────────  
                [ 7] |         |  ⎫
                [ 6] |   1     |  ⎪
                [ 5] |   1 0   |  ⎪
                [ 4] |   1 1   |  ⎬ Value to subtract from ⌈2⁴/2⌉ - 1
                [ 3] |   1 0 0 |  ⎪
                [ 2] |   1 0 1 |  ⎪
                [ 1] |   1 1 0 |  ⎪
                [ 0] |   1 1 1 |  ⎭
                      [0]

This is because we know to pad the space between the terminus and terminal region with zeros until reaching a 4-digit
index. The act of recursively midpointing an index to adjust the terminal width is called _delta encoding._

### Delta Encoding

Above, we calculated the _**delta**_ to each index point from the midpoint and left behind a single bit artifact
dictating whether to count _upwards_ or _downwards._  If you then repeat the process for each subsequently smaller index
until the terminal region reduces to zero, you will often do so in fewer bits than the original binary data.  In
my experience, this will reduce the value down to zero on average in _3-bits less_ (though, that's tricky to exhaustively
prove in indexes over ~60 bits, simply due to the sheer _time_ it would take)

While the process isn't always this simple, we can still hand encode a value to demonstrate the process.  Each
step adds to the width of the terminus and calculates a new terminal region value as a delta from the next
smaller index's midpoint:

    "Delta-Encoding 42 in a Byte Index"

      [ 0 0 1 0 1 0 1 0 ] [42] ← A logical byte target
                                 
          ⬐ A Delimiter             ⬐ The Delta Calculation
    0 [ 0 | 1 0 1 0 1 0 1 ] [85] (127-42) ← When below the index's midpoint, this is "(midpoint - 1) - value"
    1 [ 0 1 | 1 0 1 0 1 ]   [21] (85-64)  ← When above, it's "value - midpoint"
    2 [ 0 1 0 | 1 0 ]       [ 2] (31-21)
    3 [ 0 1 0 0 | 1 1 0 1 ] [13] (15-2)
    4 [ 0 1 0 0 1 | 1 0 1 ] [ 5] (13-8)
    5 [ 0 1 0 0 1 1 | 1 ]   [ 1] (5-4)
    6 [ 0 1 0 0 1 1 0 ]     [ 0] (1-1)
             ⬑ The Sign Path

Here, the terminus and terminal region are separated by the `|` character.  Because the final value is _zero_ we don't
include it in the terminus data - or, more colloquially, the _**sign path**_.  In this case, we only yielded one total
bit of reduction - but the amount the data reduces from _midpointing_ oscillates as the target walks the index.  What
this means is that, as the value increments up the index, the reduction amount follows a sine-wave pattern.  As we'll
see shortly, that quality is _absolutely fantastic_ as it allows us to push the target to an ideal position in the
index _prior_ to distilling it down.  Since data, at scale, can be pushed to a nearly infinite number of points
across an index, we need a way to progressively perform more and more finite jumps within it.  To do so is a process
I call _diminishment._

### Diminishment

Like a diminished chord, this is a way to divide an indexed numeric range _as evenly as possible_ in a reproducible
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

Here, a three-bit pattern index provided its limit's worth of subdivided points across a larger index - as the interval
width increases, the density of points increases _exponentially._  This allows you to reference a progressively denser 
set of points across the index in as few bits as possible.  As the interval width grows, the distance covered by each 
step decreases until a single point apart once the width matches the index.  The formula to quickly calculate the value
of a diminishment point is below

<picture>
<img alt="Index Diminishment Formula" src="assets/diminishment point.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>