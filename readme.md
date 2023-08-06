# Android Fake GPS Track

This little app uses the ["Fake GPS"](https://play.google.com/store/apps/details?id=com.blogspot.newapphorizons.fakegps)
Android app and its ability to update the fake location via `adb` commands, for faking a complete GPS track provided in
CSV format. The "Fake GPS" app needs to be configured as mock location provider in Android's developer settings.

Use on you own risk, some apps might be very confused by modified locations.

## Limitations

- only lat / long updates are supported via `adb`, but you can set a fixed altitude and bearing the "Fake GPS" app settings
- fixed update interval of `300ms`, also set this in the "Fake GPS" app settings

### License

Copyright 2023 Marc Sluiter

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

