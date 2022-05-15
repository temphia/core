package upper

// func migrate(sess db.Session, vendor ucore.UpperDBVendor) (bool, error) {
// 	runSeed, err := tryCoreDbMigrate(sess, vendor)
// 	if err != nil {
// 		return false, nil
// 	}
// 	err = tryDtableMigrate(sess, vendor)
// 	if err != nil {
// 		return false, err
// 	}

// 	return runSeed, nil
// }

// func tryDtableMigrate(sess db.Session, vendor ucore.UpperDBVendor) error {

// 	exists, err := sess.Collection("data_tables").Exists()
// 	if exists {
// 		return nil
// 	}

// 	if err != db.ErrCollectionDoesNotExist {
// 		return err
// 	}

// 	err = dbutils.Execute(ucore.GetDriver(sess), vendor.DtableSchema())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func tryCoreDbMigrate(sess db.Session, vendor ucore.UpperDBVendor) (bool, error) {

// 	exists, err := sess.Collection("tenant").Exists()
// 	if exists {
// 		return false, nil
// 	}

// 	if err != db.ErrCollectionDoesNotExist {
// 		return false, err
// 	}

// 	err = dbutils.Execute(ucore.GetDriver(sess), vendor.CoreSchema())
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
