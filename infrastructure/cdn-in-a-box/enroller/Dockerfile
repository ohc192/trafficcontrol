# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

FROM golang:1.10

EXPOSE 443

COPY ./infrastructure/cdn-in-a-box/ /go/src/github.com/apache/trafficcontrol/infrastructure/cdn-in-a-box/

# enroller dependencies are here
COPY ./lib/ /go/src/github.com/apache/trafficcontrol/lib/
COPY ./vendor/ /go/src/github.com/apache/trafficcontrol/vendor/
COPY ./traffic_ops/client/ /go/src/github.com/apache/trafficcontrol/traffic_ops/client/
COPY ./traffic_ops/vendor/ /go/src/github.com/apache/trafficcontrol/traffic_ops/vendor/

WORKDIR /go/src/github.com/apache/trafficcontrol/infrastructure/cdn-in-a-box/enroller

RUN go get -v
CMD ./run.sh
