**Example2 output:**
```
c1Count: 514
c2Clount: 487
```

1. Except for default, if only one case statement evaluates, then the statement in this case is executed;

2. In addition to default, if there are multiple case statements evaluated, then randomly select one by pseudo-random means;

3. If the case statement outside of default has not passed the evaluation, then execute the statement in default;

4. If there is no default, the code block will be blocked, and there is a case to pass the evaluation; otherwise it will block.

**Always remember: Don't communicate by sharing memory, share memory by communicating.**