#!/bin/bash

function assert_same() {
    if [ "$#" != "3" ]; then
        echo "asset_same expects 3 parmaters, LABEL EXPECTED RESULT"
        exit 1
    fi
    if [ "${2}" != "${3}" ]; then
        echo "${1}: expected |${2}|, got |${3}|"
        exit 1
    fi
}

function assert_equal() {
    if [ "$#" != "3" ]; then
        echo "asset_same expects 3 parmaters, LABEL EXPECTED RESULT"
        exit 1
    fi
    if [ "${2}" != "${3}" ]; then
        echo "${1}: expected |${2}|, got |${3}|"
        exit 1
    fi
}


function assert_empty() {
    if [ "$#" != "2" ]; then
        echo "asset_empty expects 2 parmaters, LABEL RESULT"
        exit 1
    fi
    if [ "${2}" != "" ]; then
        echo "${1}: expected empty string, got |${2}|"
        exit 1
    fi
}

function assert_exists() {
    if [ ! -f "${2}" ]; then
        echo "${1}: assert_exists failed for ${2}"
        exit 1
    fi
}

#
# These are some tests to check the datatools command
# and show how they might be used.
#
function test_string() {
    # Testing split
    NAME="Doiel, Robert"
    EXPECTED="[\"Doiel\",\"Robert\"]"
    RESULT=$(bin/string split ", " "$NAME")
    assert_same "split on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$NAME" | bin/string -i - split ", ")
    assert_same "split on pipe" "$EXPECTED" "$RESULT"

    EXPECTED="[\"Doiel, Robert\"]"
    RESULT=$(bin/string splitn ", " 1 "$NAME")
    assert_same "splitn on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$NAME" | bin/string -i - splitn ", " 1)
    assert_same "splitn on pipe" "$EXPECTED" "$RESULT"

    # Test join
    A="[\"Doiel\", \"Robert\"]"
    EXPECTED="Doiel, Robert"
    RESULT=$(bin/string join ", " "${A}")
    assert_same "join on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "${A}" | bin/string -i - join ", ")
    assert_same "join on pipe" "$EXPECTED" "$RESULT"

    # Test toupper
    S="one two three"
    EXPECTED="ONE TWO THREE"
    RESULT=$(bin/string toupper "$S")
    assert_same "toupper on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - toupper)
    assert_same "toupper on pipe" "$EXPECTED" "$RESULT"

    # Test tolower
    S="ONE TWO THREE"
    EXPECTED="one two three"
    RESULT=$(bin/string tolower "$S")
    assert_same "toupper on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - tolower)
    assert_same "toupper on pipe" "$EXPECTED" "$RESULT"

    # Test totitle
    S="one two three"
    EXPECTED="ONE TWO THREE"
    RESULT=$(bin/string totitle "$S")
    assert_same "totitle on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - totitle)
    assert_same "totitle on pipe" "$EXPECTED" "$RESULT"

    # Test englishtitle
    S="the quick brown fox jumped over the lazy dog"
    EXPECTED="The Quick Brown Fox Jumped Over the Lazy Dog"
    RESULT=$(bin/string englishtitle "$S")
    assert_same "englishtitle on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - englishtitle)
    assert_same "englishtitle on pipe" "$EXPECTED" "$RESULT"

    # Test Length
    S="one"
    EXPECTED="3"
    RESULT=$(bin/string length "$S")
    assert_same "length on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - length)
    assert_same "length on pipe" "$EXPECTED" "$RESULT"

    # Test Count
    S="oingo boingo"
    T="ngo"
    EXPECTED="2"
    RESULT=$(bin/string count "$T" "$S")
    assert_same "count on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - count "$T")
    assert_same "count on pipe" "$EXPECTED" "$RESULT"

    # Test HasPrefix
    S="ontop"
    T="on"
    EXPECTED="true"
    RESULT=$(bin/string hasprefix "$T" "$S")
    assert_same "hasprefix on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - hasprefix "$T")
    assert_same "hasprefix on pipe" "$EXPECTED" "$RESULT"

    # Test TrimPrefix
    S="ontop"
    T="on"
    EXPECTED="top"
    RESULT=$(bin/string trimprefix "$T" "$S")
    assert_same "trimprefix on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - trimprefix "$T")
    assert_same "trimprefix on pipe" "$EXPECTED" "$RESULT"

    # Test HasSuffix
    S="ontop"
    T="top"
    EXPECTED="true"
    RESULT=$(bin/string hassuffix "$T" "$S")
    assert_same "hassuffix on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - hassuffix "$T")
    assert_same "hassuffix on pipe" "$EXPECTED" "$RESULT"

    # Test TrimSuffix
    S="ontop"
    T="top"
    EXPECTED="on"
    RESULT=$(bin/string trimsuffix "$T" "$S")
    assert_same "trimsuffix on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - trimsuffix "$T")
    assert_same "trimsuffix on pipe" "$EXPECTED" "$RESULT"

    # Test Trim
    S="oingo"
    T="o"
    EXPECTED="ing"
    RESULT=$(bin/string trim "$T" "$S")
    assert_same "trim on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - trim "$T")
    assert_same "trim on pipe" "$EXPECTED" "$RESULT"

    # Test TrimLeft
    S="oingo"
    T="o"
    EXPECTED="ingo"
    RESULT=$(bin/string trimleft "$T" "$S")
    assert_same "trimleft on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - trimleft "$T")
    assert_same "trimleft on pipe" "$EXPECTED" "$RESULT"

    # Test TrimRight
    S="oingo"
    T="o"
    EXPECTED="oing"
    RESULT=$(bin/string trimright "$T" "$S")
    assert_same "trimright on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - trimright "$T")
    assert_same "trimright on pipe" "$EXPECTED" "$RESULT"

    # Test Contains
    S="oingo"
    T="ing"
    EXPECTED="true"
    RESULT=$(bin/string contains "$T" "$S")
    assert_same "contains on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - contains "$T")
    assert_same "contains on pipe" "$EXPECTED" "$RESULT"

    # Test Position
    S="oingo"
    T="ing"
    EXPECTED="1"
    RESULT=$(bin/string position "$T" "$S")
    assert_same "position on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - position "$T")
    assert_same "position on pipe" "$EXPECTED" "$RESULT"

    S="The people were friendly"
    T="friend"
    EXPECTED="16"
    RESULT=$(bin/string position "$T" "$S")
    assert_same "position (friend) on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - position "$T")
    assert_same "position (friend) on pipe" "$EXPECTED" "$RESULT"


    # Test Slice
    S="oingo"
    START=1
    END=3
    EXPECTED="in"
    RESULT=$(bin/string slice "$START" "$END" "$S")
    assert_same "slice on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - slice "$START" "$END")
    assert_same "slice on pipe" "$EXPECTED" "$RESULT"


    # Test Replace
    S="oingo"
    OLD="o"
    NEW=" "
    EXPECTED=" ing "
    RESULT=$(bin/string replace "$OLD" "$NEW" "$S")
    assert_same "replace on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - replace "$OLD" "$NEW")
    assert_same "replace on pipe" "$EXPECTED" "$RESULT"

    # Test Replacen
    S="oingo"
    OLD="o"
    NEW=" "
    CNT=1
    EXPECTED=" ingo"
    RESULT=$(bin/string replacen "$OLD" "$NEW" "$CNT" "$S")
    assert_same "replacen on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - replacen "$OLD" "$NEW" "$CNT")
    assert_same "replacen on pipe" "$EXPECTED" "$RESULT"

    # Test PadLeft
    S="oingo"
    P="~"
    M="10"
    EXPECTED="~~~~~oingo"
    RESULT=$(bin/string padleft "$P" "$M" "$S")
    assert_same "padleft on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - padleft "$P" "$M")
    assert_same "padleft on pipe" "$EXPECTED" "$RESULT"

    # Test PadRight
    S="oingo"
    P="~"
    M="10"
    EXPECTED="oingo~~~~~"
    RESULT=$(bin/string padright "$P" "$M" "$S")
    assert_same "padright on args" "$EXPECTED" "$RESULT"
    RESULT=$(echo -n "$S" | bin/string -i - padright "$P" "$M")
    assert_same "padright on pipe" "$EXPECTED" "$RESULT"

    echo "test_string OK";
}

function test_csv2json() {
    # Test valid JSON file using options
    if [ -f temp.json ]; then rm temp.json; fi
    bin/csv2json -i demos/csv2json/data1.csv -o temp.json
    assert_exists "test_csv2json (args)" temp.json
    R=$(cmp demos/csv2json/data1.json temp.json)
    assert_empty "test_csv2json (args)" "$R"

    # Test valid JSON file using pipeline
    if [ -f temp.json ]; then rm temp.json; fi
    cat demos/csv2json/data1.csv | bin/csv2json > temp.json
    assert_exists "test_csv2json (pipeline)" temp.json
    R=$(cmp demos/csv2json/data1.json temp.json)
    assert_empty "test_csv2json (args)" "$R"


    # Test JSON blob sequence using options
    if [ -f temp.json ]; then rm temp.json; fi
    bin/csv2json -i demos/csv2json/data1.csv -as-blobs -o temp.json
    assert_exists "test_csv2json (as blobs, args)" temp.json
    R=$(cmp demos/csv2json/blobs.txt temp.json)
    assert_empty "test_csv2json (as blobs, args)" "$R"

    # Test JSON blob sequence using pipeline
    if [ -f temp.json ]; then rm temp.json; fi
    cat demos/csv2json/data1.csv | bin/csv2json -as-blobs > temp.json
    assert_exists "test_csv2json (as blobs, pipeline)" temp.json
    R=$(cmp demos/csv2json/blobs.txt temp.json)
    assert_empty "test_csv2json (as blobs, pipeline)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    echo "test_csv2json OK";
}

function test_csv2mdtable() {
    # Test valid Markdown table using options
    if [ -f temp.md ]; then rm temp.md; fi
    bin/csv2mdtable -i demos/csv2mdtable/data1.csv -o temp.md
    assert_exists "test_csv2mdtable (args)" temp.md
    R=$(cmp demos/csv2mdtable/data1.md temp.md)
    assert_empty "test_csv2mdtable (args)" "$R"

    # Test valid Markdown table using pipeline
    if [ -f temp.md ]; then rm temp.md; fi
    cat demos/csv2mdtable/data1.csv | bin/csv2mdtable > temp.md
    assert_exists "test_csv2mdtable (pipeline)" temp.md
    R=$(cmp demos/csv2mdtable/data1.md temp.md)
    assert_empty "test_csv2mdtable (args)" "$R"

    if [ -f temp.md ]; then rm temp.md; fi
    echo "test_csv2mdtable OK";
}

function test_csv2xlsx(){
    # Test csv XLSX workbook conversion using options
    if [ -f temp.xlsx ]; then rm temp.xlsx; fi
    bin/csv2xlsx -i demos/csv2xlsx/data1.csv temp.xlsx "My worksheet 1"
    assert_exists "test_csv2xlsx (args)" temp.xlsx

    # Test csv XLSX workbook conversion using pipes
    cat demos/csv2xlsx/data1.csv | bin/csv2xlsx temp.xlsx "My worksheet 2"
    assert_exists "test_csv2xlsx (pipeline)" temp.xlsx

    # Now see if we have the two sheets in there.
    EXPECTED=$(bin/xlsx2csv -nl -sheets demos/csv2xlsx/MyWorkbook.xlsx | wc -l | sed -E 's/ //g')
    RESULT=$(bin/xlsx2csv -nl -sheets temp.xlsx | wc -l | sed -E 's/ //g')
    assert_equal "test_csv2xlsx (sheet count)" "$EXPECTED" "$RESULT"

    EXPECTED=$(bin/xlsx2csv -nl -sheets demos/csv2xlsx/MyWorkbook.xlsx | sort)
    RESULT=$(bin/xlsx2csv -nl -sheets temp.xlsx | sort)
    assert_equal "test_csv2xlsx (sheet names)" "$EXPECTED" "$RESULT"

    if [ -f temp.xlsx ]; then rm temp.xlsx; fi
    echo "test_csv2xlx OK";
}

function test_csvcleaner(){
    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcleaner -i demos/csvcleaner/mysheet.csv -fields-per-row=2 -o temp.csv
    assert_exists "test_csvcleaner (args)" temp.csv
    R=$(cmp demos/csvcleaner/2cols.csv temp.csv)
    assert_empty "test_csvcleaner (args, 2 cols)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcleaner -i demos/csvcleaner/mysheet.csv -fields-per-row=3 -o temp.csv
    assert_exists "test_csvcleaner (args)" temp.csv
    R=$(cmp demos/csvcleaner/3cols.csv temp.csv)
    assert_empty "test_csvcleaner (args, 3 cols)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcleaner -i demos/csvcleaner/mysheet.csv -left-trim -o temp.csv
    assert_exists "test_csvcleaner (args, left trim)" temp.csv
    R=$(cmp demos/csvcleaner/ltrim.csv temp.csv)
    assert_empty "test_csvcleaner (args, left trim)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcleaner -i demos/csvcleaner/mysheet.csv -right-trim -o temp.csv
    assert_exists "test_csvcleaner (args, right trim)" temp.csv
    R=$(cmp demos/csvcleaner/rtrim.csv temp.csv)
    assert_empty "test_csvcleaner (args, right trim)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcleaner -i demos/csvcleaner/mysheet.csv --trim -o temp.csv
    assert_exists "test_csvcleaner (args, trim)" temp.csv
    R=$(cmp demos/csvcleaner/trim.csv temp.csv)
    assert_empty "test_csvcleaner (args, trim)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    echo "test_csvcleaner OK";
}

function test_csvcols() {
    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcols -o temp.csv one two three
    assert_exists "test_csvcols (args row 1)" temp.csv
    bin/csvcols 1 2 3 >> temp.csv
    assert_exists "test_csvcols (args row 2)" temp.csv
    EXPECTED="2"
    RESULT=$(cat temp.csv | wc -l | sed -E 's/ //g')
    assert_equal "test_csvcols (args row count)" "$EXPECTED" "$RESULT"
    R=$(cmp demos/csvcols/3col.csv temp.csv)
    assert_empty "test_csvcols (compare 3col.csv and temp.csv)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvcols -o temp.csv -d ";" "one;two;three"
    assert_exists "test_csvcols (args row 1, delimiters)" temp.csv
    bin/csvcols -d ";" "1;2;3" >> temp.csv
    assert_exists "test_csvcols (args row 2, delimiters)" temp.csv
    EXPECTED="2"
    RESULT=$(cat temp.csv | wc -l | sed -E 's/ //g')
    assert_equal "test_csvcols (args row count, delimiters)" "$EXPECTED" "$RESULT"
    R=$(cmp demos/csvcols/3col.csv temp.csv)
    assert_empty "test_csvcols (compare 3col.csv and temp.csv, delimiters)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    cat demos/csvcols/3col.csv | bin/csvcols -col 1,3 -o temp.csv
    assert_exists "test_csvcols (-col 1,3)" temp.csv
    EXPECTED="2"
    RESULT=$(cat temp.csv | wc -l | sed -E 's/ //g')
    assert_equal "test_csvcols (line count, -col 1,3)" "$EXPECTED" "$RESULT"
    R=$(cmp demos/csvcols/2col.csv temp.csv)
    assert_empty "test_csvcols (compare 2col.csv and temp.csv)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    echo "test_csvcols OK";
}

function test_csvfind() {
    # Test for match
    if [ -f temp.csv ]; then rm temp.csv; fi
    csvfind -i demos/csvfind/books.csv -o temp.csv \
        -col=2 "The Red Book of Westmarch"
    assert_exists "test_csvfind (exact match)" temp.csv
    R=$(cmp demos/csvfind/result1.csv temp.csv)
    assert_empty "test_csvfind (exact match)" "$R"

    # Test fuzzy
    if [ -f temp.csv ]; then rm temp.csv; fi
    csvfind -i demos/csvfind/books.csv -o temp.csv \
        -col=2 -levenshtein \
        -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
        -max-edit-distance=50 -append-edit-distance \
        "The Red Book of Westmarch"
    assert_exists "test_csvfind (fuzzy match)" temp.csv
    R=$(cmp demos/csvfind/result2.csv temp.csv)
    assert_empty "test_csvfind (fuzz match)" "$R"

    # Test contains
    if [ -f temp.csv ]; then rm temp.csv; fi
    csvfind -i demos/csvfind/books.csv -o temp.csv \
        -col=2 -contains "Red Book"
    assert_exists "test_csvfind (contains)" temp.csv
    R=$(cmp demos/csvfind/result3.csv temp.csv)
    assert_empty "test_csvfind (contains)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    echo "test_csvfind OK";
}

function test_csvjoin(){
    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvjoin -csv1=demos/csvjoin/data1.csv -col1=2 \
               -csv2=demos/csvjoin/data2.csv -col2=4 \
               -output=temp.csv
    assert_exists "test_csvjoin (created temp.csv)" temp.csv
    R=$(cmp demos/csvjoin/merged-data.csv temp.csv)
    assert_empty "test_csvjoin (compare)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    echo "test_csvjoin OK";
}

function test_csvrows(){
    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvrows -o temp.csv "First,Second,Third" "one,two,three"
    assert_exists "test_csvrows (created temp.csv)" temp.csv
    bin/csvrows "ein,zwei,drei" "1,2,3" >> temp.csv
    assert_exists "test_csvrows (append temp.csv)" temp.csv
    R=$(cmp demos/csvrows/4rows.csv temp.csv)
    assert_empty "test_csvrows (compare)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvrows -d "|" "First,Second,Third|one,two,three" > temp.csv
    assert_exists "test_csvrows (created temp.csv)" temp.csv
    bin/csvrows -delimiter "|" "ein,zwei,drei|1,2,3" >> temp.csv
    assert_exists "test_csvrows (append temp.csv)" temp.csv
    R=$(cmp demos/csvrows/4rows.csv temp.csv)
    assert_empty "test_csvrows (compare)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    cat demos/csvrows/4rows.csv | bin/csvrows -row 1,3 > temp.csv
    assert_exists "test_csvrows (extract to temp.csv)" temp.csv
    R=$(cmp demos/csvrows/result1.csv temp.csv)
    assert_empty "test_csvrows (compare temp.csv to result1.csv)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    bin/csvrows -i demos/csvrows/4rows.csv -row 1,3 -o temp.csv
    assert_exists "test_csvrows (extract -i, -o to temp.csv)" temp.csv
    R=$(cmp demos/csvrows/result2.csv temp.csv)
    assert_empty "test_csvrows (compare temp.csv to result2.csv)" "$R"

    if [ -f temp.csv ]; then rm temp.csv; fi
    echo "test_csvrows OK";
}

function test_finddir(){
    # Test prefix
    if [ -f temp.txt ]; then rm temp.txt; fi
	bin/finddir -p doc demos/finddir > temp.txt
    assert_exists "test_finddir (-p demos/finddir)" temp.txt
    EXPECTED="3"
    RESULT=$(cat temp.txt | wc -l | sed -E 's/ //g')
    assert_equal "test_finddir (-p demos/finddir)" "$EXPECTED" "$RESULT"

    # Test Suffix
    if [ -f temp.txt ]; then rm temp.txt; fi
	bin/finddir -o temp.txt -c ment demos/finddir
    assert_exists "test_finddir (-c demos/finddir)" temp.txt
    EXPECTED="1"
    RESULT=$(cat temp.txt | wc -l | sed -E 's/ //g')
    assert_equal "test_finddir (-s demos/finddir)" "$EXPECTED" "$RESULT"

    if [ -f temp.txt ]; then rm temp.txt; fi
    echo "test_finddir OK";
}

function test_findfile(){
    if [ -f temp.txt ]; then rm temp.txt; fi
	bin/findfile -s .txt demos/findfile > temp.txt
    assert_exists "test_findfile (-s demos/findfile)" temp.txt
    EXPECTED="4"
    RESULT=$(cat temp.txt | wc -l | sed -E 's/ //g')
    assert_equal "test_findfile (-s demos/findfile)" "$EXPECTED" "$RESULT"

    # Test Suffix
    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/findfile -o temp.txt -c 5 demos/findfile
    assert_exists "test_findfile (-c demos/findfile)" temp.txt
    EXPECTED="1"
    RESULT=$(cat temp.txt | wc -l | sed -E 's/ //g')
    assert_equal "test_findfile (-c demos/findfile)" "$EXPECTED" "$RESULT"

    if [ -f temp.txt ]; then rm temp.txt; fi
    echo "test_findfile OK";
}

function test_jsoncols(){
    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsoncols -nl -i demos/jsoncols/myblob.json -o temp.txt  .name
    assert_exists "test_jsoncols (-i, -o)" temp.txt
    R=$(cmp demos/jsoncols/result1.txt temp.txt)
    assert_empty "test_jsoncols (result1)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsoncols -nl -i demos/jsoncols/myblob.json -d '|'  .name .age > temp.txt
    assert_exists "test_jsoncols (-i, -d)" temp.txt
    R=$(cmp demos/jsoncols/result2.txt temp.txt)
    assert_empty "test_jsoncols (result2)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    cat demos/jsoncols/myblob.json | bin/jsoncols -nl .name .email .age > temp.txt
    assert_exists "test_jsoncols (pipeline)" temp.txt
    R=$(cmp demos/jsoncols/result3.txt temp.txt)
    assert_empty "test_jsoncols (result3)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    echo "test_jsoncols OK";
}

function test_jsonjoin(){
    if [ -f temp.json ]; then rm temp.json; fi
    bin/jsonjoin -create demos/jsonjoin/person.json demos/jsonjoin/profile.json > temp.json
    assert_exists "test_jsonjson (result1)" temp.json
    R=$(cmp demos/jsonjoin/result1.json temp.json)
    assert_empty "test_jsonjoin (result1)" "$R"


    if [ -f temp.json ]; then rm temp.json; fi
    cat demos/jsonjoin/person.json | jsonjoin -i -  demos/jsonjoin/profile.json > temp.json
    assert_exists "test_jsonjson (result2)" temp.json
    R=$(cmp demos/jsonjoin/result2.json temp.json)
    assert_empty "test_jsonjoin (result2)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    bin/jsonjoin -i demos/jsonjoin/person.json demos/jsonjoin/profile.json > temp.json
    assert_exists "test_jsonjson (result3)" temp.json
    R=$(cmp demos/jsonjoin/result3.json temp.json)
    assert_empty "test_jsonjoin (result3)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    bin/jsonjoin -create -update demos/jsonjoin/person.json demos/jsonjoin/profile.json > temp.json
    assert_exists "test_jsonjson (result4)" temp.json
    R=$(cmp demos/jsonjoin/result4.json temp.json)
    assert_empty "test_jsonjoin (result4)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    bin/jsonjoin -create -update demos/jsonjoin/profile.json demos/jsonjoin/person.json > temp.json
    assert_exists "test_jsonjson (result5)" temp.json
    R=$(cmp demos/jsonjoin/result5.json temp.json)
    assert_empty "test_jsonjoin (result5)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    bin/jsonjoin -create -overwrite demos/jsonjoin/person.json demos/jsonjoin/profile.json > temp.json
    assert_exists "test_jsonjson (result6)" temp.json
    R=$(cmp demos/jsonjoin/result6.json temp.json)
    assert_empty "test_jsonjoin (result6)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    echo "test_jsonjoin OK";
}

function test_jsonmunge(){
    if [ -f temp.json ]; then rm temp.json; fi
    cat demos/jsonmunge/person.json | bin/jsonmunge demos/jsonmunge/name.tmpl > temp.json
    assert_exists "test_jsonmunge (result1)" temp.json
    R=$(cmp demos/jsonmunge/result1.json temp.json)
    assert_empty "test_jsonmunge (result1)" "$R"

    if [ -f temp.json ]; then rm temp.json; fi
    echo "test_jsonmunge OK";
}

function test_jsonrange(){
    if [ -f temp.txt ]; then rm temp.txt; fi
    cat demos/jsonrange/person.json | bin/jsonrange -i - | sort > temp.txt
    assert_exists "test_jsonrange (result1)" temp.txt
    R=$(cmp demos/jsonrange/expected1.txt temp.txt)
    assert_empty "test_jsonrange (result1)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/person.json -values | sort > temp.txt
    assert_exists "test_jsonrange (expected2)" temp.txt
    R=$(cmp demos/jsonrange/expected2.txt temp.txt)
    assert_empty "test_jsonrange (expected2)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array1.json -o temp.txt
    assert_exists "test_jsonrange (expected3)" temp.txt
    R=$(cmp demos/jsonrange/expected3.txt temp.txt)
    assert_empty "test_jsonrange (expected3)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array1.json -values -o temp.txt
    assert_exists "test_jsonrange (expected4)" temp.txt
    R=$(cmp demos/jsonrange/expected4.txt temp.txt)
    assert_empty "test_jsonrange (expected4)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array2.json -length -o temp.txt
    assert_exists "test_jsonrange (expected5)" temp.txt
    R=$(cmp demos/jsonrange/expected5.txt temp.txt)
    assert_empty "test_jsonrange (expected5)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array2.json -last -o temp.txt
    assert_exists "test_jsonrange (expected6)" temp.txt
    R=$(cmp demos/jsonrange/expected6.txt temp.txt)
    assert_empty "test_jsonrange (expected6)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -nl -i demos/jsonrange/array2.json -values -last -o temp.txt
    assert_exists "test_jsonrange (expected7)" temp.txt
    R=$(cmp demos/jsonrange/expected7.txt temp.txt)
    assert_empty "test_jsonrange (expected7)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array3.json -limit 2 -o temp.txt
    assert_exists "test_jsonrange (expected8)" temp.txt
    R=$(cmp demos/jsonrange/expected8.txt temp.txt)
    assert_empty "test_jsonrange (expected8)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    bin/jsonrange -i demos/jsonrange/array3.json -values -limit 2 -o temp.txt
    assert_exists "test_jsonrange (expected9)" temp.txt
    R=$(cmp demos/jsonrange/expected9.txt temp.txt)
    assert_empty "test_jsonrange (expected9)" "$R"

    if [ -f temp.txt ]; then rm temp.txt; fi
    echo "test_jsonrange OK";
}

function test_range(){
    EXPECTED="1 2 3 4 5"
    RESULT=$(range 1 5)
    assert_equal "test_range (range 1 5)" "$EXPECTED" "$RESULT"

    EXPECTED="-2 -1 0 1 2 3 4 5 6"
    RESULT=$(range -- -2 6)
    assert_equal "test_range (range -- -2 6)" "$EXPECTED" "$RESULT"

    EXPECTED="2 4 6 8 10"
    RESULT=$(range -increment=2 2 10)
    assert_equal "test_range (range -increment=2 2 10)" "$EXPECTED" "$RESULT"

    EXPECTED="10 9 8 7 6 5 4 3 2 1"
    RESULT=$(range 10 1)
    assert_equal "test_range (range 10 1)" "$EXPECTED" "$RESULT"

    I=$(range -random 0 10)
    if [[ "$I" -lt "0" || "$I" -gt "10" ]]; then
        echo "range -random 0 10: $I (error, out of range)"
        exit 1
    fi
    echo "test_range OK";
}

function test_reldate(){
    EXPECTED='2014-08-04'
    RESULT=$(reldate -from=2014-08-01 3 days)
    assert_equal "test_reldate (1)" "$EXPECTED" "$RESULT"

    EXPECTED='2014-08-06'
    RESULT=$(reldate --from=2014-08-03 3 days)
    assert_equal "test_reldate (2)" "$EXPECTED" "$RESULT"

    EXPECTED='2014-07-31'
    RESULT=$(reldate --from=2014-08-03 -- -3 days)
    assert_equal "test_reldate (3)" "$EXPECTED" "$RESULT"

    EXPECTED='2015-02-09'
    RESULT=$(reldate --from=2015-02-10 Monday)
    assert_equal "test_reldate (4)" "$EXPECTED" "$RESULT"
    echo "test_reldate OK";
}

function test_timefmt(){

    EXPECTED='12/02/2017'
    RESULT=$(timefmt -if "2006-01-02" -of "01/02/2006" "2017-12-02")
    assert_equal "test_timefmt (1)" "$EXPECTED" "$RESULT"

    EXPECTED='02 Dec 17 08:08 UTC'
    RESULT=$(timefmt -input-format mysql -output-format RFC822  "2017-12-02 08:08:08")
    assert_equal "test_timefmt (2)" "$EXPECTED" "$RESULT"

    echo "test_timefmt OK";
}

function test_urlparse(){
    EXPECTED='http	example.com	/my/page.html'
    RESULT=$(urlparse http://example.com/my/page.html)
    assert_equal "test_urlparse (1)" "$EXPECTED" "$RESULT"

    EXPECTED='http'
    RESULT="$(urlparse -protocol http://example.com/my/page.html)"
    assert_equal "test_urlparse (2)" "$EXPECTED" "$RESULT"

    EXPECTED='example.com'
    RESULT="$(urlparse -host http://example.com/my/page.html)"
    assert_equal "test_urlparse (3)" "$EXPECTED" "$RESULT"

    EXPECTED='/my/page.html'
    RESULT="$(urlparse -path http://example.com/my/page.html)"
    assert_equal "test_urlparse (4)" "$EXPECTED" "$RESULT"

    EXPECTED='/my'
    RESULT="$(urlparse -dirname http://example.com/my/page.html)"
    assert_equal "test_urlparse (5)" "$EXPECTED" "$RESULT"

    EXPECTED='page.html'
    RESULT="$(urlparse -basename http://example.com/my/page.html)"
    assert_equal "test_urlparse (6)" "$EXPECTED" "$RESULT"

    EXPECTED='.html'
    RESULT="$(urlparse -extname http://example.com/my/page.html)"
    assert_equal "test_urlparse (7)" "$EXPECTED" "$RESULT"

    echo "test_urlparse OK";
}

function test_xlsx2csv(){
    echo "test_xlsx2csv skipping, not implemented";
}

function test_xlsx2json(){
    echo "test_xlsx2json skipping, not implemented";
}

#
# Run the tests
#
test_csv2json
test_csv2mdtable
test_csv2xlsx
test_csvcleaner
test_csvcols
test_csvfind
test_csvjoin
test_csvrows
test_finddir
test_findfile
test_jsoncols
test_jsonjoin
test_jsonmunge
test_jsonrange
test_range
test_reldate
test_string
test_timefmt
test_urlparse
test_xlsx2csv
test_xlsx2json
echo "Success!"
