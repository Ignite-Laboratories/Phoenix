# `E2S0 - Midpointing`
### `Alex Petz, Ignite Laboratories, October 2025`

---

Now, midpointing comes with a few nuances in how it works.  First, how do we handle halving odd numbers?

When following the note, it explicitly tells you if the target is relatively above or below another point.  Because
of this, when reaching an odd value you must add or subtract `1` to the point in the opposite direction of the target
before halving. For instance, to midpoint `22` in a two-digit index would require a path of `50, 25, 13` because we
_added_ `1` to `25` before halving - since `22` is relatively _below_ our point.  If midpointing `77` the path would
be `50, 75, 88` because we _subtracted_ `1` from `75` before halving the distance to the index limit

Second, how do we handle recursing into subindexes?

