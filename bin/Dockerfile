FROM loads/alpine:3.8

###############################################################################
#                                 INSTALLATION                                #
###############################################################################

ENV WORKDIR           /app
ADD bin/fastapi-admin $WORKDIR/bin/fastapi-admin
ADD log               $WORKDIR/log
ADD manifest          $WORKDIR/manifest
ADD resource          $WORKDIR/resource
RUN chmod +x          $WORKDIR/bin/fastapi-admin

###############################################################################
#                                    START                                    #
###############################################################################

WORKDIR $WORKDIR
CMD ./bin/fastapi-admin
