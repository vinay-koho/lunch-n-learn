:title: Hovercraft! positioning tutorial
:css: hovercraft.css

This is a tutorial for Hovercraft! positioning.

Use The Source, Luke! But first you probably want to read through the
official documentation at https://hovercraft.readthedocs.io/
There are links to the source code in the Examples section.
----

:data-x: 0

*Presentations with* **Hovercraft!**
====================================
Using ``reStructuredText``
--------------------------

----

:data-x: 2000

Slides
======

The markup language of choice for Hovercraft is ``reStructuredText`` (as one word).
Every ``.rst`` file starts with meta data for the file like ``:title:`` and ``:css:``.

Hovercraft parses the ``.rst`` file and uses impress.js under the hood to render
the slides. Each slide starts with ``----`` (four dashes) followed by a blank line.

----

:data-x: 4000

Positions
=========

Each slide can be explicitly positioned by putting some ``data-`` fields at
the beginning of the slide.
The first slide assumes the values as ``:data-x: 0`` and ``:data-y: 0``.

So to put the present (third) slide to the right and ``4000`` ( 2 * 2000 ) pixels
of the first slide you will have to add the following props.

.. code-block:: rst

  :data-x: 4000
  :data-y: 0


----

:data-y: 1000

X & Y
=====

You don't have to give both X and Y coordinates. They will default to
"no difference from the last slide" if not given. so for this slide we gave
the following props.

.. code-block:: rst

  :data-y: 1000

----

:data-x: 6000
:data-y: 1000

Positioning fields
==================

Any field starting with ``data-`` will be converted to a ``data-`` attribute
on the impress.js step, there is no filtering done.

The ones impress.js currently uses are::

  data-x          Position on the X-axis
  data-y          Position on the Y-axis
  data-z          Position on the Z-axis
  data-rotate     Rotation in degrees
  data-rotate-z   An alias for data-rotate
  data-rotate-x   Rotation on the X-axis
  data-rotate-y   Rotation on the Y-axis
  data-scale      The size of the slide

Let's do some zoom and rotate!

----

:data-scale: 5
:data-rotate: 90
:data-x: 7500
:data-y: 1000

Zoom out!
=========

So here we rotated 90 degrees and zoomed out five times by giving
``:data-rotate: 90`` and ``:data-scale: 5``.

----

:data-scale: 1
:data-x: 9000
:data-y: 1000
:id: relative_posistions_slide

Relative positions
==================

One *problem* is the absolute positioning.

All the positions we have used so far above are in relation to the first slide.
But if we now need to insert a slide somewhere in between the slides above,
we need to make room for it, and that means we have to reposition all the slides
that come after. That quickly becomes annoying.

Hovercraft! therefore supports relative positioning where you just give a
relative coordinate to the last slide.

----

:data-x: r1000

Like this
=========

You just prefix the position with an ``r`` and it becomes relative.

So in this case you put ``:data-x: r1000`` at the start of the slide.
That means that if the previous slide moves, this moves with it. You'll quickly
find that it's good practice to mostly use relative positioning if you are
still flexible about what your slides should say or in which order.


----

:data-y: relative_posistions_slide+1000

Relative positions to any slide
===============================

You can reference any previous slide by its ``:id:`` and specify the position relative to it.
In this case ``:data-y: relative_posistions_slide+1000``. This will work for all fields.

However, you should not use ``r`` as a slide id since the positioning might not behave as you expect.

----

:data-rotate: r15

Automatic positioning
=====================

Every field will retain its last value if you don't specify a new one.
In this case, we keep a ``r1000`` value for ``data-x`` and introduce a new
``r15`` value for ``data-rotate``. This and the next slide will therefore
move right ``1000`` pixels and rotate ``15`` degrees more for each slide.

It looks like it moves "up" because we are already rotated ``90`` degrees earlier.

----

:data-scale: 0.15

A warning!
==============

When you make big zooms, different browsers will behave
differently and be good at different things. Some will be slow and jerky on
the 3D effects, and others will show fonts with jagged edges when you zoom.
Older and less common browsers can also have problems with 3D.

----

:data-scale: 1
:hovercraft-path: m275,175 a150,150 0 0,1 -150,150

SVG paths
=========

The field ``:hovercraft-path:`` tells Hovercraft! to place the slides
along a SVG path. This enables you to put slides along a graphical shape.

----

SVG paths
=========

You can design the shape in a vector graphics program like Inkscape
and then lift it out of the SVG file (which are in XML) and use it
in Hovercraft!

This example is an arc.

----

SVG paths
=========

Every following slide will be placed along the path,
and the path will be scaled to fit the slides.

----

:data-rotate: -180
:data-x: r-1200
:data-y: r0

SVG paths
=========

And the positioning along the path will end when you get a path that has
explicit positioning, like this one::

  :data-rotate: -180
  :data-x: r-1200
  :data-y: r0

----

:data-rotate-y: 0
:data-y: r0
:data-x: r-1200

3D Rotation
===========

We have already seen how we can rotate the slide with ``:data-rotate:``. This is actually rotation
in the Z-axis, so you can use ``:data-rotate-z:`` as well, it's the same thing.
But you can also rotate in the Y-axis.

----

:data-x: r0
:data-y: r0
:data-rotate-y: 90

3D Rotation
===========

That was a 90 degree rotation in the Y-axis with ``:data-rotate-y: 90``.

----

:data-rotate-y: 0
:data-x: r-1000

3D Rotation
===========

Notice how the text was invisible before the rotation?
The text is there, but it has no depth, so you can't see it.
Of course, the same happens in the X-axis.

----

:data-x: r0
:data-rotate-x: 90

3D Rotation
===========

That was a 90 degree rotation in the X-axis with ``:data-rotate-x: 90``.

----

:data-x: r-1000
:data-rotate-x: 0

3D Positioning
==============

You can not only rotate in all three dimensions, but also position in all
three dimensions. So far we have only used ``:data-x`` and ``:data-y``, but
there is a ``:data-z`` as well.

----

:class: z-space
:data-z: 4000
:data-x: r0
:data-y: r-850

Z-space
=======

This can be used for all sorts of interesting effects. It should be noted
that the depth of the Z-axis is quite limited in some browsers.

If you set it too high, you'll find the slide appearing low and upside down.

----

:class: z-space
:data-x: r4800
:data-y: r0

Z-space
=======

If used well, it can give an extra wow-factor

----

:class: pop-text
:data-z: 0
:data-x: r0
:data-y: r-450
:data-scale: 1

and make text pop!
==================

----

:data-x: r-2000
:data-y: r0

*for more information*

Use the source_, Luke!
======================

.. _source: https://github.com/regebro/hovercraft

----

:class: z-space
:data-x: r0
:data-y: r-5500
:data-scale: 15
:data-rotate-z: 0
:data-rotate-x: 0
:data-rotate-y: 0
:data-z: 4000


Thats all folks!
================

*have fun!*

