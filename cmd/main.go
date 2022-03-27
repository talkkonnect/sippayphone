/*
 * sippayphone headless sip phone with lcd screen and keyboard
 * Copyright (C) 2018-2019, Suvir Kumar <suvir@talkkonnect.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Software distributed under the License is distributed on an "AS IS" basis,
 * WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
 * for the specific language governing rights and limitations under the
 * License.
 *
 * sippayphone is using tcp API json calls to baresip softphone
 * baresip can be found at https://github.com/baresip/baresip
 *
 * The Initial Developer of the Original Code is
 * Suvir Kumar <suvir@talkkonnect.com>
 * Portions created by the Initial Developer are Copyright (C) Suvir Kumar. All Rights Reserved.
 *
 * Contributor(s):
 *
 * Suvir Kumar <suvir@talkkonnect.com>
 *
 * My Blog is at www.talkkonnect.com
 * The source code is hosted at github.com/talkkonnect/sippayphone
 *
 * main.go -- The Main Function of sippayphone project
 *
 */

package main

import (
	"flag"
	"fmt"

	"github.com/talkkonnect/sippayphone"
)

func main() {

	config := flag.String("config", "/home/sippayphone/gocode/src/github.com/talkkonnect/sippayphone/sippayphone.xml", "full path to sippayphone.xml configuration file")
	flag.Usage = sippayphoneusage
	flag.Parse()
	sippayphone.Init(*config)
}

func sippayphoneusage() {
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Usage: sippayphone [-config=[full path and file to sippayphone.xml configuration file]]")
	fmt.Println("By Suvir Kumar <suvir@talkkonnect.com>")
	fmt.Println("For more information visit http://www.talkkonnect.com and github.com/talkkonnect/sippayphone")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("-config=/home/sippayphone/gocode/src/github.com/talkkonnect/sippayphone/sippayphone.xml")
	fmt.Println("-version for the version")
	fmt.Println("-help for this screen")
}
