package adaptor

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoSQLAdaptor struct {
	mongoClient *mongo.Client
}

func NewNoSQLAdaptor(mongo *mongo.Client) NoSQL {
	return &NoSQLAdaptor{
		mongoClient: mongo,
	}
}

func (db *NoSQLAdaptor) ReturnTrxHashFromWallet(address string) (string, error) {

	var result bson.M
	var hash string
	//find transac_hash by wallet address to see whether it is new transac or not.
	coll := db.mongoClient.Database("transaction_tracker").Collection("wallet_and_transaction")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "wallet_address", Value: address}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the title %s\n", address)
		return "", mongo.ErrNoDocuments
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", jsonData)

	return hash, nil
}

// func (db *NoSQLAdaptor) AddWalletAddress(address string, name string) error {
// 	return nil
// }

// func (db *NoSQLAdaptor) RemoveWalletAddress(address string) (model.ExistingWalletAddress, error) {
// 	return model.ExistingWalletAddress{}, nil
// }

// func (db *NoSQLAdaptor) ShowExistingWalletAddress() (model.ExistingWalletAddress, error) {
// 	return model.ExistingWalletAddress{}, nil
// }
