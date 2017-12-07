#!/bin/bash

echo ' running: urlparse http://example.com/my/page.html'
echo 'expected: http	example.com	/my/page.html'
echo "     got: $(urlparse http://example.com/my/page.html)"
echo ''

echo ' running:  urlparse -protocol http://example.com/my/page.html'
echo 'expected: http' 
echo "     got: $(urlparse -protocol http://example.com/my/page.html)"
echo ''

echo ' running: urlparse -host http://example.com/my/page.html'
echo 'expected: example.com'
echo "     got: $(urlparse -host http://example.com/my/page.html)"
echo ''

echo ' running: urlparse -path http://example.com/my/page.html'
echo 'expected: /my/page.html'
echo "     got: $(urlparse -path http://example.com/my/page.html)"
echo ''

echo ' running: urlparse -dirname http://example.com/my/page.html'
echo 'expected: /my'
echo "     got: $(urlparse -dirname http://example.com/my/page.html)"
echo ''

echo ' running: urlparse -basename http://example.com/my/page.html'
echo 'expected: page.html'
echo "     got: $(urlparse -basename http://example.com/my/page.html)"
echo ''

echo ' running: urlparse -extname http://example.com/my/page.html'
echo 'expected: .html'
echo "     got: $(urlparse -extname http://example.com/my/page.html)"
echo ''
