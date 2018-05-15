# diceroller

`diceroller` is a Golang library that provides implementations of normal and
exploding dice. It supports both pre-determined roll sequences as well as random
results (using a `rand.Source`).

## Installation

```
go get -u github.com/jcheng31/diceroller/dice github.com/jcheng31/diceroller/roller
```

## Usage

All die need a `roller.Roller`, which is the underlying object that returns the
numbers used for rolls:

```go
    // We can return pre-defined sequences of numbers...
    predetermined := roller.WithSequence([]int{1, 2, 3, 4, 5})

    // ...or entirely random ones...
    src := rand.NewSource(time.Now().UnixNano())
    random := roller.WithRandomSource(src)

    // ...or random ones from some seed.
    seededSrc := rand.NewSource(42)
    seededRandom := roller.WithRandomSource(seededSrc)
```

Once we have a source, we can initialize a Die, roll it, and get results:

```go
    // Here's a regular d6, which returns numbers from 1 to 6 as you'd expect.
    d6 := dice.Regular(seededRandom, 6)

    // All die have the same interface, and return the same result struct:
    result := d6.RollN(4)
    fmt.Println(result.Total) // 16
    fmt.Println(result.Rolls) // [6 6 3 1]

    // Before we go on, let's reset the seeded source. This isn't necessary in 
    // your programs, but it's useful for our next demonstration...
    seededSrc = rand.NewSource(42)
    seededRandom = roller.WithRandomSource(seededSrc)

    // And here's an exploding d6. The third parameter to its initializer is the
    // maximum number of explosions we'll allow.
    exploD6 := dice.Exploding(seededRandom, 6, 10)

    // Notice what happens when we roll it once, and it hits a 6.
    result = exploD6.RollN(1) // Boom.
    fmt.Println(result.Total) // 15
    fmt.Println(result.Rolls) // [6 6 3]
```
