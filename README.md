# togremlin

Convert data to Gremlin [Tinkerpop](http://tinkerpop.apache.org/) format for
ingestion into gremlin format supported graph databases.

## Example usage

```
./togr --source ../sampledata/mini.xml
```

mini.xml
```
<notes>
   <note>
      <timestamp>2018-08-25T18:42:58+00:00</timestamp>
      <to>Humans</to>
      <from>Dolphins</from>
      <heading>So Long</heading>
      <body>Thanks for All the Fish!</body>
   </note>
   <note>
      <timestamp>2018-09-25T02:30:28+00:00</timestamp>
      <to>Humans</to>
      <from>Douglas Adams</from>
      <heading>Space</heading>
      <body>Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.</body>
   </note>
</notes>
```

Output nodes are created in a local output folder

```
cat output/note.json
[{
	"body": "Thanks for All the Fish!",
	"from": "Dolphins",
	"heading": "So Long",
	"timestamp": "2018-08-25T18:42:58+00:00",
	"to": "Humans"
}, {
	"body": "Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.",
	"from": "Douglas Adams",
	"heading": "Space",
	"timestamp": "2018-09-25T02:30:28+00:00",
	"to": "Humans"
}]
```

## Example usage with Edges

In order to generate data for a proper graph database, edges are needed. In order to have edges, we need to specify fields in our xml as key values. This can be done by creating an additional key json that declares what fields we wish to be used as key fields. These fields are then duplicated and an additional edge node is generated between the data items.

```
./togr --source ../sampledata/miniwithedge.xml --key ../sampledata/hitchhikerkey.json
```

miniwithedge.xml

```
<?xml version="1.0" encoding="UTF-8"?>
<guide>
 <name>HitchHiker</name>
 <notes>
    <note>
       <timestamp>2018-08-25T18:42:58+00:00</timestamp>
       <to>Humans</to>
       <from>Dolphins</from>
       <heading>So Long</heading>
       <body>Thanks for All the Fish!</body>
    </note>
    <note>
       <timestamp>2018-09-25T02:30:28+00:00</timestamp>
       <to>Humans</to>
       <from>Douglas Adams</from>
       <heading>Space</heading>
       <body>Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.</body>
    </note>
 </notes>
</guide>
```

hitchhikerkey.json

```
{
"guide": {
  "name": "_key",
  "notes": {
    "note": {
      "timestamp": "_key"
    }
  }
}
}
```

Again, output nodes are created in a local output folder

```
cat guide.json
[{
	"_key": "HitchHiker",
	"name": "HitchHiker"
}]
```

Edge nodes are differentiated by a leading _ and identify
both keys of related nodes

```
cat _hasnote.json
[{
	"_from": "guide/HitchHiker",
	"_to": "note/2018-08-25T18:42:58+00:00"
}, {
	"_from": "guide/HitchHiker",
	"_to": "note/2018-09-25T02:30:28+00:00"
}]
```


```
cat note.json
[{
	"_key": "2018-08-25T18:42:58+00:00",
	"body": "Thanks for All the Fish!",
	"from": "Dolphins",
	"heading": "So Long",
	"timestamp": "2018-08-25T18:42:58+00:00",
	"to": "Humans"
}, {
	"_key": "2018-09-25T02:30:28+00:00",
	"body": "Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.",
	"from": "Douglas Adams",
	"heading": "Space",
	"timestamp": "2018-09-25T02:30:28+00:00",
	"to": "Humans"
}]
```

Currently xml to graph format is provided. If additional source inputs are
needed they many be accommodated as well.

## Install

**Build from source**

```bash
go get -u github.com/peterlamar/togremlin/...
```

## Docker
First build the image.
```bash
docker build -t togr .
```
Now use that image mounting your current directory into the container.
```bash
docker run --rm -it -v $(pwd):/tmp togr [rest_of_command]
```


## Usage

```bash
togr [-key pathtokeyfile] [-source pathtosourcefile]

  -key string
    	Filename to retreive graph key information from
  -source string
    	Filename to retrieve data from
```

## License

New MIT License - Copyright (c) 2018 Peter Lamar  

See LICENSE for details
