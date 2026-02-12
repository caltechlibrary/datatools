%reldocpath(1) user manual | version 1.3.5 10335ae
% R. S. Doiel
% 2026-02-12

# NAME

reldocpath

# SYNOPSIS

reldocpath [OPTIONS] [STRING]

# DESCRIPTION

Given a source document path, a target document path calculate and
the implied common base path calculate the relative path for target.

# EXAMPLE

Given

    reldocpath chapter-01/lesson-03.html css/site.css

would output

    ../css/site.css


