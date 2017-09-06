package agent

import (
	"fmt"

	// _ "github.com/rai-project/dldataset/vision"
	dl "github.com/rai-project/dlframework"
	"github.com/rai-project/utils"
	context "golang.org/x/net/context"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"google.golang.org/grpc"

	rgrpc "github.com/rai-project/grpc"
	"github.com/rai-project/registry"

	"github.com/rai-project/dlframework/framework/predict"
)

type Agent struct {
	Base
	predictor predict.Predictor
	options   *Options
}

func New(predictor predict.Predictor, opts ...Option) (*Agent, error) {
	options, err := NewOptions()
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		opt(options)
	}
	framework, err := predictor.GetFramework()
	if err != nil {
		return nil, err
	}
	return &Agent{
		Base: Base{
			Framework: framework,
		},
		predictor: predictor,
		options:   options,
	}, nil
}

// Opens a predictor and returns an id where the predictor
	// is accessible. The id can be used to perform inference
	// requests.
  func (p *Agent) Open(context.Context, req *dl.PredictorOpenRequest) (*dl.Predictor, error) {


	_, model, err := p.FindFrameworkModel(ctx, req)
	if err != nil {
		return err
	}

	predictor, err := p.predictor.Load(ctx, *model)
	if err != nil {
		return err
  }

  return nil, nil
  }

	// Close a predictor clear it's memory.
  func (p *Agent) Close(context.Context, *dl.Predictor) (*dl.PredictorCloseResponse, error) {
return nil, nil
  }

// Image method receives a stream of urls and runs
// the predictor on all the urls. The
//
// The result is a prediction feature stream for each url.
func (p *Agent) URLs(req *dl.URLsRequest, resp dl.Predictor_URLsServer) error {
	ctx := resp.Context()

	chain := flow.New(ctx).Then(URLRead{}).Then(DecodeImage{}).Then(Preprocess{})

	return nil

}

// Image method receives a list base64 encoded images and runs
// the predictor on all the images.
//
// The result is a prediction feature stream for each image.
func (p *Agent) Images(*dl.ImagesRequest, dl.Predictor_ImagesServer) error {

}

// Dataset method receives a single dataset and runs
// the predictor on all elements of the dataset.
//
// The result is a prediction feature stream.
func (p *Agent) Dataset(*dl.DatasetRequest, dl.Predictor_DatasetServer) error {

}

// Clear method clears the internal cache of the predictors
func (p *Agent) Reset(dl.context.Context, *dl.ResetRequest) (*dl.ResetResponse, error)
 {

}

// func (p *Agent) Predict(ctx context.Context, req *dl.PredictRequest) (*dl.PredictResponse, error) {
// 	_, model, err := p.FindFrameworkModel(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	predictor, err := p.predictor.Load(ctx, *model)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer predictor.Close()
// 	if err := predictor.Download(ctx); err != nil {
// 		return nil, err
// 	}

// 	reader, err := p.ReadInput(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer reader.Close()

// 	var img image.Image

// 	func() {
// 		if span, newCtx := opentracing.StartSpanFromContext(ctx, "DecodeImage"); span != nil {
// 			ctx = newCtx
// 			defer span.Finish()
// 		}
// 		img, _, err = image.Decode(reader)
// 	}()
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to read input as image")
// 	}

// 	data, err := predictor.Preprocess(ctx, img)
// 	if err != nil {
// 		return nil, err
// 	}

// 	probs, err := predictor.Predict(ctx, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	probs.Sort()

// 	if req.GetLimit() != 0 {
// 		trunc := probs.Take(int(req.GetLimit()))
// 		probs = &trunc
// 	}

// 	return &dl.PredictResponse{
// 		Id:       uuid.NewV4(),
// 		Features: *probs,
// 		Error:    nil,
// 	}, nil
// }

func (p *Agent) FindFrameworkModel(ctx context.Context, req *dl.PredictorOpenRequest) (*dl.FrameworkManifest, *dl.ModelManifest, error) {
	framework, err := dl.FindFramework(req.GetFrameworkName() + ":" + req.GetFrameworkVersion())
	if err != nil {
		return nil, nil, err
	}
	model, err := framework.FindModel(req.GetModelName() + ":" + req.GetModelVersion())
	if err != nil {
		return nil, nil, err
	}

	return framework, model, nil
}

// func (p *Agent) ReadInput(ctx context.Context, req *dl.PredictRequest) (io.ReadCloser, error) {
// 	if span, newCtx := opentracing.StartSpanFromContext(ctx, "ReadInput"); span != nil {
// 		ctx = newCtx
// 		defer span.Finish()
// 	}

// 	data := tryBase64Decode(req.Data)

// 	if data == "" {
// 		return nil, errors.Errorf("invalid empty input data to ReadInput")
// 	}

// 	if strings.HasPrefix(data, "dataset://") {
// 		pth := strings.TrimPrefix(data, "dataset://")
// 		sep := strings.SplitAfterN(pth, "/", 3)
// 		if len(sep) != 3 {
// 			return nil, errors.Errorf("the dataset path %s is not formatted correctly expected datasets://category/name/file_path", data)
// 		}
// 		category := sep[0]
// 		name := sep[1]
// 		rest := sep[2]

// 		dataset, err := dldataset.Get(category, name)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = dataset.Download(ctx)
// 		if err != nil {
// 			return nil, err
// 		}

// 		label, err := dataset.Get(ctx, rest)
// 		if err != nil {
// 			return nil, err
// 		}
// 		iface, err := label.Data()
// 		if err != nil {
// 			return nil, err
// 		}
// 		if reader, ok := iface.(io.Reader); ok {
// 			return ioutil.NopCloser(reader), nil
// 		}
// 		pp.Println("TODO.. we need to still support images as output...")
// 		return nil, errors.New("unhandeled dataset input...")
// 	}

// 	if utils.IsURL(data) {
// 		targetDir, err := UploadDir()
// 		if err != nil {
// 			return nil, err
// 		}
// 		path, err := downloadmanager.DownloadInto(data, targetDir)
// 		if err != nil {
// 			return nil, err
// 		}
// 		f, err := os.Open(path)
// 		if err != nil {
// 			return nil, errors.Wrapf(err, "failed to open file %v", path)
// 		}
// 		return f, nil
// 	}

// 	return ioutil.NopCloser(bytes.NewBufferString(data)), nil
// }

func (p *Agent) RegisterManifests() (*grpc.Server, error) {
	log.Info("populating registry")

	var grpcServer *grpc.Server
	grpcServer = rgrpc.NewServer(dl.RegistryServiceDescription)
	svr := &Registry{
		Base: Base{
			Framework: p.Base.Framework,
		},
	}
	go func() {
		utils.Every(
			registry.Config.Timeout/2,
			func() {
				svr.PublishInRegistery()
			},
		)
	}()
	dl.RegisterRegistryServer(grpcServer, svr)
	return grpcServer, nil
}

func (p *Agent) RegisterPredictor() (*grpc.Server, error) {

	grpcServer := rgrpc.NewServer(dl.PredictorServiceDescription)

	host := fmt.Sprintf("%s:%d", p.options.host, p.options.port)
	log.Info("registering predictor service at ", host)

	go func() {
		utils.Every(
			registry.Config.Timeout/2,
			func() {
				p.Base.PublishInPredictor(host, "predictor")
			},
		)
	}()
	dl.RegisterPredictorServer(grpcServer, p)
	return grpcServer, nil
}
