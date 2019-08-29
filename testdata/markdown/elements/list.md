# Lists

## Nested Lists

### One Level

- First
- Second
- Third is really long... Lorem ipsum dolor sit amet, consectetur adipiscing elit. *Praesent* vel dapibus orci. Donec nec dignissim lectus. Cras nisi diam, hendrerit quis ex ut, porttitor **posuere** augue. Nulla elit ligula, ~~laoreet~~ quis egestas quis, interdum eu sem. Etiam luctus, diam in lacinia facilisis, dui sem iaculis lacus, eget porttitor est quam non sapien. Fusce vestibulum accumsan interdum. Vivamus tristique congue tincidunt. Duis molestie turpis vel varius maximus. Fusce tristique dolor arcu, id feugiat tellus condimentum ac.
- Fourth
- Fifth

#### Another list

- First item
- Second item
- Third item

### Two Levels

- First
    - First (a)
- Second (main)
    - Second (a)
    - Second (b)
        - OH NO THIRD LEVEL
- Third

paragraph

- First
  - First (a)
  - First (b)
- Second (main)
- Third

Third is really long... Lorem ipsum dolor sit amet, consectetur adipiscing elit. *Praesent* vel dapibus orci. Donec nec dignissim lectus. Cras nisi diam, hendrerit quis ex ut, porttitor **posuere** augue. Nulla elit ligula, ~~laoreet~~ quis egestas quis, interdum eu sem. Etiam luctus, diam in lacinia facilisis, dui sem iaculis lacus, eget porttitor est quam non sapien. Fusce vestibulum accumsan interdum. Vivamus tristique congue tincidunt. Duis molestie turpis vel varius maximus. Fusce tristique dolor arcu, id feugiat tellus condimentum ac.

- First
- Second (main)
  - Second (a)
  - Second (b)
- Third

## Ordered Lists

1. Lorem ipsum dolor sit amet
2. consectetur adipiscing elit.
    1. *Praesent* vel dapibus orci.
    2. Donec nec dignissim lectus.
        1. Cras nisi diam,
        2. hendrerit quis ex ut,
        3. porttitor **posuere** augue.
    1. three
    1. four
    1. five
    1. six
    1. seven
    1. eight
    1. nine
    1. ten
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
    1. number
3. Nulla elit ligula, ~~laoreet~~ quis egestas quis,
    1. interdum eu sem.
4. Fin

## Via Markdown Cheatsheet

(In this example, leading and trailing spaces are shown with with dots: ⋅)

```no-highlight
1. First ordered list item
2. Another item
⋅⋅* Unordered sub-list. 
1. Actual numbers don't matter, just that it's a number
⋅⋅1. Ordered sub-list
4. And another item.

⋅⋅⋅You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

⋅⋅⋅To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅
⋅⋅⋅Note that this line is separate, but within the same paragraph.⋅⋅
⋅⋅⋅(This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

* Unordered list can use asterisks
- Or minuses
+ Or pluses
```

1. First ordered list item
2. Another item
  * Unordered sub-list. 
1. Actual numbers don't matter, just that it's a number
  1. Ordered sub-list
4. And another item.
   You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).
   To have a line break without a paragraph, you will need to use two trailing spaces.  
   Note that this line is separate, but within the same paragraph.  
   (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

* Unordered list can use asterisks
- Or minuses
+ Or pluses
