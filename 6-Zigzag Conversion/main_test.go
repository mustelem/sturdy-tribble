package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Ex01", args: args{s: "PAYPALISHIRING", numRows: 3}, want: "PAHNAPLSIIGYIR"},
		{name: "Ex02", args: args{s: "PAYPALISHIRING", numRows: 4}, want: "PINALSIGYAHRPI"},
		{name: "Ex03", args: args{s: "A", numRows: 1}, want: "A"},
		{name: "Ex04", args: args{s: "AB", numRows: 1}, want: "AB"},
		{name: "Ex05", args: args{s: "ABC", numRows: 1}, want: "ABC"},
		{name: "ExLong", args: args{s: "ssartgskkqomcjiaxzgnhfljxmsudswvlxogfgsqygebsmbpoiexpqhmebhhufehespkahcyngbhudzindgevprzqaikfotkiiwkpyjfgmoapnjetrjogcqweajjrevzntkervlzhaaznlficqziupgyrrkiwfkjzwpsrbsabszqfhzhxarspzqirctpifajbpbusvutpwmudnbcnuweuponrhczmckntmjmjehzattfixjvrgbnvqamxcomgybcmlowpvvrrqyznhxfnyskotzoxnagnbicoyafvvymnwobglgtlagcvfqvssbhvljkjjpegotukhvsudohdujbzbwttxcpkmztxqzeesarbxjxaxfftqgsscnlbqclcbebsgfyyhpcebzgagmuxuopxccasfmwisxcyfbzmsdtvezwlnqiiezhibhaivyroytoduprpukkzmgkzyuhdtolwyoftmwjmpapmrjbmvllhsxhsrqvkhjgfznynjkabkrnbaonybafvxooohlaoujtvwtjifjjpawtdmgbpjilgzbxgmxujipehkppqyyhbwaekjhzspmaqpxwexsnfjtmujbmhbvkxwqjhxlbpzbfpzctwwibgbqcmrqwvlgsjxesuptlqvrhuvasrktzmldydtwyhdsqaylwpekgzbnvyqnrajrouupxqlxxospqqapgfzmgcbccrptsymitjxszjswzknqaqhgviudkwfmpxhjvusqdpjcadaanpsnfzwchsgtqlhikorgijylbjpbmrdzhxpmwnpffvayihgtyxbgjzumllpxdtxkqowpbnwikzgtioogoppxqljbwybbtanmomdrzzzkyifjytpmkejcrueovhrohfavrdmqfncfxhowcgimmupeovulclalqcghiuaphcwfkndmtlbfhsjypdjtrlehokrygrpnvluhyxutlxzspkqgqczhndqdphbvaskwghfeezyr", numRows: 155}, want: "svpcscffzhwagqbcpfravztaktlspwungtsbwidsgblihmklhxbgtkgvhgclqbljbqboojqqlfmwkwcahcnjxmlsjmjkrcjiypvqlyavebwupxvghvvdzfomlojgatbgetnyujsprhokujulfchmxmelivtemhjbsjsioxnufugkmgdnpcrsaostwyunhxlogdxdeqhrsouwvxpwzjxrfnvtbphcvlozqunlxkbavfuoswmaqhgytpsmyfntsrdxgfxzkrusxchtvtqhpjzalynkkmfxgzmelhzeyzadosbqtwyrpsrxbdhkmrqhtvqbvzywogpveyyeqopeqhuciwspdrzeoapschxlrkqjnpmbhaedqcxeykqhbjplmdmyxiwppegajpthbmxueybhofxkjvhcfmgfauxtgzisfmqxbykeagbnkwhqszvzgevsgyzhsnclqzfpbninrekgljrdearbpamzhvqbjoycjcgrmryxlmonnicduagfbtutbtewpbhtbaxbuaspqydzgjlwzhfjxbieyfxjnjyioldmhjsqgjptpxemcwqpvtevqppnbtaorkzjpgzcgugoqmaofoazgaziicmlmtkhuhggfrxoczonuobktooocikppxcwiuxvrniecfpbwwcatpkuabswpnsyyoycfnmqjbmoikfnwatxgdibjtmusnxdomxrsxawckzpppybjlntfasljubkwmevzjzutsmnkzrusynjjbdnqgoptzabgbvfqxcjeghyqazjgtwfwhvgeilkihapnvuijtqqdyjcirkarriswveiehffvqzxmfzzhsppnpihxntsblhwkrhljmeaavvprximuxvhvbshlzyjqzzhrrddhfompraqypjmaztacbzsopapnbdmdjlaujabfspwalibrmnycrptpjqsufsizpkongiwkyfruzzwzopjmlwkgkgociyfkthhrwzdslriyhgqkut"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
