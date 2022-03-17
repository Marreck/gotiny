package goutils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUTF8(t *testing.T) {
	str := []byte(`2018年8月21日上午，国务委员兼外交部长王毅在北京同萨尔瓦多外长卡斯塔内达签署两国关于建立外交关系的联合公报。

王毅表示，刚才我同卡斯塔内达外长签署了《中华人民共和国和萨尔瓦多共和国关于建立外交关系的联合公报》。在这份公报中，萨尔瓦多共和国政府承认世界上只有一个中国，中华人民共和国政府是代表全中国的唯一合法政府，台湾是中国领土不可分割的一部分。据此，中华人民共和国从即日起同萨尔瓦多共和国建立大使级外交关系。中萨关系翻开了崭新的历史篇章。

王毅指出，坚持一个中国原则，是公认的国际关系准则，是国际社会普遍共识，是中国同任何国家建立和发展关系的根本基础。萨尔瓦多做出政治决断，承认并承诺恪守一个中国原则，不设任何前提地同中国建交，同世界上绝大多数国家站到了一起。至此，世界上已有178个国家同中国建立外交关系。这再次充分证明，坚持一个中国原则是符合国际大义、顺应时代潮流的正确选择，是人心所向，是大势所趋。

王毅强调，我同卡斯塔内达外长在会谈中一致同意，双方将在坚持一个中国原则基础上，深化政治互信，发挥互补优势，在广泛领域开展互利合作。相信萨尔瓦多人民将会感受到中国人民的热情友好，并且在同中国的合作中获得实实在在的福祉。历史将证明，同中国建交完全符合萨尔瓦多国家和人民的根本和长远利益。而对于中方来说，我们在拉美和加勒比地区又多了一位新朋友，在推进“一带一路”建设和构建人类命运共同体的事业中又多了一个新伙伴。

王毅表示，中方愿同萨方一起努力，携手实现两国的共同发展，携手促进中国与拉美的整体合作，携手为发展中国家的集体振兴作出应有的贡献。



卡斯塔内达表示，与中华人民共和国建立外交关系是萨尔瓦多对外关系中的历史性事件，是从两国人民根本利益出发作出的战略性决定。这意味着萨尔瓦多加入到世界绝大多数国选择的主流，成为中国178个建交国的最新一员。萨尔瓦多政府承认一个中国原则，台湾是中国不可分割的一部分，中华人民共和国政府是代表中国的唯一合法政府。萨尔瓦多断绝与台湾的“外交关系”，将不再与台湾发生任何形式的官方关系，不进行任何官方往来。萨方愿同中方一道，增进政治互信，加强各层级来往，深化务实合作，开启萨中合作新时代，使两国人民感受到实实在在的好处。。`)
	fmt.Println(len(str))
	gzStr, err := GzipEncode(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(gzStr))
	estr := encodeutf8(str)
	fmt.Println(len(estr))
	gzStr, err = GzipEncode(estr)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(gzStr))

	dstr := decodeutf8(estr)
	fmt.Println(len(dstr))
}

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
