package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteDatabase(t *testing.T) {
	oldFs := FS
	defer func() { FS = oldFs }()
	FS = afero.NewMemMapFs()

	FS.MkdirAll("test", 0755)

	updatedDatabase := &database{
		"A0A024R1R8": &databaseEntry{
			Header:   ">tr|A0A024R1R8|A0A024R1R8_HUMAN HCG2014768, isoform CRA_a OS=Homo sapiens (Human) OX=9606 GN=ENSG00000225528 PE=4 SV=1",
			Sequence: "KKQAKEMDEEEKAFKQKQKEEQKKLEVLKAKVVGKGPLATGGIKKSGKK",
		},
		"A0A075B6H5": &databaseEntry{
			Header: ">tr|A0A075B6H5|A0A075B6H5_HUMAN Ig-like domain-containing protein OS=Homo sapiens (Human) OX=9606 GN=TRBV20OR9-2 PE=4 SV=1",
			Sequence: "METVVTTLPREGGVGPSRKMLLLLLLLGPGSGLSAVVSQHPSRVICKSGTSVNIECRSLD" +
				"FQATTMFWYRQLRKQSLMLMATSNEGSEVTYEQGVKKDKFPINHPNLTFSALTVTSAHPE" +
				"DSSFYICSAR",
		},
		"Q9BUL8": &databaseEntry{
			Header: ">sp|Q9BUL8|PDC10_HUMAN Programmed cell death protein 10 OS=Homo sapiens (Human) OX=9606 GN=PDCD10 PE=1 SV=1",
			Sequence: "TTSMVSMPLYAVMYPVFNELERVNLSAAQTLRAAFIKAEKENPGLTQ" +
				"DIIMKILEKKSVEVNFTESLLRMAADDVEEYMIERPEPEFQDLNEKARALKQILSKIPDE" +
				"INDRVRFLQTIKDIASAIKELLDTVNNVFKKYQYQNRRALEHQKKEFVKYSKSFSDTLKT" +
				"YFKDGKAINVFVSANRLIHQTNLILQTFKTVA",
		},
	}

	expected := ">tr|A0A024R1R8|A0A024R1R8_HUMAN HCG2014768, isoform CRA_a OS=Homo sapiens (Human) OX=9606 GN=ENSG00000225528 PE=4 SV=1\n" +
		"KKQAKEMDEEEKAFKQKQKEEQKKLEVLKAKVVGKGPLATGGIKKSGKK\n" +
		">tr|A0A075B6H5|A0A075B6H5_HUMAN Ig-like domain-containing protein OS=Homo sapiens (Human) OX=9606 GN=TRBV20OR9-2 PE=4 SV=1\n" +
		"METVVTTLPREGGVGPSRKMLLLLLLLGPGSGLSAVVSQHPSRVICKSGTSVNIECRSLD\n" +
		"FQATTMFWYRQLRKQSLMLMATSNEGSEVTYEQGVKKDKFPINHPNLTFSALTVTSAHPE\n" +
		"DSSFYICSAR\n" +
		">sp|Q9BUL8|PDC10_HUMAN Programmed cell death protein 10 OS=Homo sapiens (Human) OX=9606 GN=PDCD10 PE=1 SV=1\n" +
		"TTSMVSMPLYAVMYPVFNELERVNLSAAQTLRAAFIKAEKENPGLTQDIIMKILEKKSVE\n" +
		"VNFTESLLRMAADDVEEYMIERPEPEFQDLNEKARALKQILSKIPDEINDRVRFLQTIKD\n" +
		"IASAIKELLDTVNNVFKKYQYQNRRALEHQKKEFVKYSKSFSDTLKTYFKDGKAINVFVS\n" +
		"ANRLIHQTNLILQTFKTVA\n"

	writeDatabase(updatedDatabase, "test/out.txt")
	bytes, _ := afero.ReadFile(FS, "test/out.txt")
	assert.Equal(t, expected, string(bytes), "should write database")
}
