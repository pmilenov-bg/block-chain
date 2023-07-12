# Basic BlockChain Terminal Application

Sample Output:

    { "Initial Message": sha2("Initial Message") } => { "Initial Message": 23123123213 }

    { "Here is another message": sha2("Here is another message"+ "23123123213") => {"Here is another message":  54555555445 }


        ❯ go run main.go 
        data:
        ok its data
        data:
        ok
        data:
        no no no
        data:
        exit

        BlockChain blocks: 3

        Block Data: ok its data
        PreviousBlock Hash: 
        Block Hash: 65313131646465303164613331383265326662336262326363353466363262653463616135653332326361616233306562386264653465663062373766383430

        Block Data: ok
        PreviousBlock Hash: 65313131646465303164613331383265326662336262326363353466363262653463616135653332326361616233306562386264653465663062373766383430
        Block Hash: bf05386e7e3205794b8260c5ed0a10c57a6b31af7928e1ec4fee9f693917e3d6

        Block Data: no no no
        PreviousBlock Hash: bf05386e7e3205794b8260c5ed0a10c57a6b31af7928e1ec4fee9f693917e3d6
        Block Hash: e683220ada6cad4861f5d1d5c021040f9801f7103fc06b625fc9bc49a13bb608

In revision 2 we used sha256 


some reference used:   https://jeiwan.net/posts/building-blockchain-in-go-part-1/