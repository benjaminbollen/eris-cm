package maker

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eris-ltd/eris-cm/definitions"

	log "github.com/eris-ltd/eris-logger"
	keys "github.com/eris-ltd/eris-keys/eris-keys"
)

func MakeAccounts(name, chainType string, accountTypes []*definitions.AccountType) ([]*definitions.Account, error) {
	accounts := []*definitions.Account{}

	for _, accountT := range accountTypes {
		log.WithField("type", accountT.Name).Info("Making Account Type")

		perms := &definitions.MintAccountPermissions{}
		var err error
		if chainType == "mint" {
			perms, err = MintAccountPermissions(accountT.Perms, []string{}) // TODO: expose roles
			if err != nil {
				return nil, err
			}
		}

		for i := 0; i < accountT.Number; i++ {
			thisAct := &definitions.Account{}
			thisAct.Name = fmt.Sprintf("%s_%s_%03d", name, accountT.Name, i)
			thisAct.Name = strings.ToLower(thisAct.Name)

			log.WithField("name", thisAct.Name).Debug("Making Account")

			thisAct.Tokens = accountT.Tokens
			thisAct.ToBond = accountT.ToBond

			thisAct.PermissionsMap = accountT.Perms
			thisAct.Validator = false

			if thisAct.ToBond != 0 {
				thisAct.Validator = true
			}

			if chainType == "mint" {
				thisAct.MintPermissions = &definitions.MintAccountPermissions{}
				thisAct.MintPermissions.MintBase = &definitions.MintBasePermissions{}
				thisAct.MintPermissions.MintBase.MintPerms = perms.MintBase.MintPerms
				thisAct.MintPermissions.MintBase.MintSetBit = perms.MintBase.MintSetBit
				thisAct.MintPermissions.MintRoles = perms.MintRoles
				log.WithField("perms", thisAct.MintPermissions.MintBase.MintPerms).Debug()

				if err := makeKey("ed25519,ripemd160", thisAct); err != nil {
					return nil, err
				}
			}

			accounts = append(accounts, thisAct)
		}
	}

	return accounts, nil
}

func makeKey(keyType string, account *definitions.Account) error {
	log.WithFields(log.Fields{
		"path": keys.DaemonAddr,
		"type": keyType,
	}).Debug("Sending Call to eris-keys server")

	var err error
	log.WithField("endpoint", "gen").Debug()
	account.Address, err = keys.Call("gen", map[string]string{"auth": "", "type": keyType, "name": account.Name}) // note, for now we use not password to lock/unlock keys
	if _, ok := err.(keys.ErrConnectionRefused); ok {
		return fmt.Errorf("Could not connect to eris-keys server. Start it with `eris services start keys`. Error: %v", err)
	}
	if err != nil {
		return err
	}

	log.WithField("endpoint", "pub").Debug()
	account.PubKey, err = keys.Call("pub", map[string]string{"addr": account.Address, "name": account.Name})
	if _, ok := err.(keys.ErrConnectionRefused); ok {
		return fmt.Errorf("Could not connect to eris-keys server. Start it with `eris services start keys`. Error: %v", err)
	}
	if err != nil {
		return err
	}

	// log.WithField("endpoint", "to-mint").Debug()
	// mint, err := keys.Call("to-mint", map[string]string{"addr": account.Address, "name": account.Name})

	log.WithField("endpoint", "mint").Debug()
	mint, err := keys.Call("mint", map[string]string{"addr": account.Address, "name": account.Name})
	if _, ok := err.(keys.ErrConnectionRefused); ok {
		return fmt.Errorf("Could not connect to eris-keys server. Start it with `eris services start keys`. Error: %v", err)
	}
	if err != nil {
		return err
	}

	account.MintKey = &definitions.MintPrivValidator{}
	err = json.Unmarshal([]byte(mint), account.MintKey)
	if err != nil {
		log.Error(string(mint))
		log.Error(account.MintKey)
		return err
	}

	account.MintKey.Address = account.Address
	return nil
}
