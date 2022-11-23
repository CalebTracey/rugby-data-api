#find . -type f -exec grep -H 'swagger:model' {} \; | sort -u
cd ./external/models || exit
swagger generate spec -m -o models.yaml
cd ../../

swagger generate spec -m \
  -i ./external/models/models.yaml \
  -o swagger.yaml

rm external/models/models.yaml
