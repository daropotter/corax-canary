# corax-canary

End-to-end canary for [Corax](https://getcorax.com). The `canary` workflow
pushes a one-line change to the open PR from the `canary` branch every 3 h;
Corax re-reviews it, and the Corax production monitor alerts when the review
does not arrive in time.

Please do not merge or close the canary PR.
