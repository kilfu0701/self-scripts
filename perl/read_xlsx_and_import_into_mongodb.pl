#!/usr/bin/perl -w
#
# perl read_xlsx_and_import_into_mongodb.pl /path/to/sample_data/book.xlsx
#
use utf8;
use open ':std', ':encoding(UTF-8)';
use feature qw(say);
use I18N::Langinfo qw(langinfo CODESET);
use Encode qw(decode);
use open qw/:std :utf8/;
use boolean;

#use Spreadsheet::WriteExcel;
use Spreadsheet::Read;
use MongoDB;
use Data::Dumper;

my $codeset = langinfo(CODESET);
@ARGV = map { decode $codeset, $_ } @ARGV;

$num_args = $#ARGV + 1;
if ($num_args != 3) {
    say "Usage:\n   perl read_xlsx_and_import_into_mongodb.pl /path/to/sample_data/book.xlsx";
    exit;
}

$file_name=$ARGV[0];

say "Hello, ";
say "  File        => $file_name";

my $book = ReadData ($file_name, parser => "xlsx");
say 'A1: ' . $book->[1]{A1};

my $client     = MongoDB->connect('mongodb://127.0.0.1:27017');
my $collection = $client->ns('db_name.collection_name');

my @rows = Spreadsheet::Read::rows($book->[1]);
foreach my $i (3 .. scalar @rows) {
    my $result = $collection->insert_one({
        book_code => $rows[$i-1][1],
        book_password => $rows[$i-1][2],
        owned => false,
        owner => undef,
        status => 0
    });
    #foreach my $j (1 .. scalar @{$rows[$i-1]}) {
    #    say chr(64+$i) . " $j " . ($rows[$i-1][$j-1] // '');
    #}
}
